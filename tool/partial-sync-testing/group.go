package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"berty.tech/go-orbit-db/accesscontroller"
	"berty.tech/go-orbit-db/iface"
	"berty.tech/go-orbit-db/stores"
	"berty.tech/weshnet/pkg/tinder"

	"github.com/libp2p/go-libp2p/core/event"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"go.uber.org/zap"
)

type group struct {
	name    string
	peers   []*peer
	mocknet mocknet.Mocknet
	msrv    *tinder.MockDriverServer
	logger  *zap.Logger
}

func createNewGroup(ctx context.Context, peersAmount int, name, storeDir string, logger *zap.Logger) *group {
	if peersAmount < 0 {
		panic("peersAmount can't be less than 0")
	}

	mn := mocknet.New()
	msrv := tinder.NewMockDriverServer()
	peers := make([]*peer, peersAmount)
	identities := []string{}

	for i := range peers {
		peers[i] = createNewPeer(ctx, fmt.Sprintf("Peer%d", i), storeDir, mn, msrv, logger)
		identities = append(identities, peers[i].odb.Identity().ID)
	}

	storeType := "eventlog"
	create := true

	for i := range peers {
		options := &iface.CreateDBOptions{
			AccessController: accesscontroller.NewEmptyManifestParams(),
			StoreType:        &storeType,
			Create:           &create,
		}

		options.AccessController.SetAccess("write", identities)
		address, err := peers[i].odb.DetermineAddress(ctx, name, storeType, &iface.DetermineAddressOptions{AccessController: options.AccessController})
		if err != nil {
			panic(err)
		}

		peers[i].store, err = peers[i].odb.Open(ctx, address.String(), options)

		if err := peers[i].eventHandler(ctx,
			func(ctx context.Context, p *peer, sub event.Subscription) {
				defer sub.Close()

				var evt interface{}
				for {
					select {
					case <-ctx.Done():
						return
					case evt = <-sub.Out():
						logEvent(p.logger, evt)
					}
				}
			}); err != nil {
			panic(err)
		}
	}

	return &group{
		name:    name,
		peers:   peers,
		mocknet: mn,
		msrv:    msrv,
		logger:  logger.Named(name),
	}
}

func (g *group) connectPeers(a, b int) error {
	if a < 0 || a >= len(g.peers) {
		return errors.New("index 'a' is out of bound")
	}
	if b < 0 || b >= len(g.peers) {
		return errors.New("index 'b' is out of bound")
	}
	if a == b {
		return errors.New("index 'a' and 'b' should not be identical")
	}

	aID := g.peers[a].ipfs.API().ID()
	bID := g.peers[b].ipfs.API().ID()

	if _, err := g.mocknet.LinkPeers(aID, bID); err != nil {
		return err
	}
	if _, err := g.mocknet.ConnectPeers(aID, bID); err != nil {
		return err
	}

	return nil
}

func (g *group) disconnectPeers(a, b int) error {
	if a < 0 || a >= len(g.peers) {
		return errors.New("index 'a' is out of bound")
	}
	if b < 0 || b >= len(g.peers) {
		return errors.New("index 'b' is out of bound")
	}
	if a == b {
		return errors.New("index 'a' and 'b' should not be identical")
	}

	aID := g.peers[a].ipfs.API().ID()
	bID := g.peers[b].ipfs.API().ID()

	if err := g.mocknet.DisconnectPeers(aID, bID); err != nil {
		return err
	}

	return g.mocknet.UnlinkPeers(aID, bID)
}

func (g *group) createGroupBranch(ctx context.Context, timeoutMS int, name string, msgAmount int, outOfOrder bool, peerIndexes []int) error {
	var (
		wgStart     sync.WaitGroup
		wgEnd       sync.WaitGroup
		errsMutex   sync.Mutex
		tCtx        context.Context
		cancel      context.CancelFunc
		peersAmount = len(peerIndexes)
		errs        = make([]error, peersAmount)
	)

	if timeoutMS > 0 {
		tCtx, cancel = context.WithTimeout(ctx, time.Duration(timeoutMS)*time.Millisecond)
	} else if msgAmount == 0 {
		// Force set a timeout as a stop condition to avoid dead lock
		tCtx, cancel = context.WithTimeout(ctx, time.Duration(peersAmount)*5*time.Second)
	} else {
		tCtx, cancel = context.WithCancel(ctx)
	}

	wgStart.Add(peersAmount)
	wgEnd.Add(peersAmount)

	// Disconnect peers on exit
	defer func() {
		cancel()
		for _, i := range peerIndexes {
			for _, j := range peerIndexes {
				if i == j {
					continue
				}
				_ = g.disconnectPeers(i, j)
			}
		}
	}()

	// Connect peers to each other
	for _, i := range peerIndexes {
		for _, j := range peerIndexes {
			if i == j {
				continue
			}
			if err := g.connectPeers(i, j); err != nil {
				return err
			}
		}
	}

	for e, i := range peerIndexes {
		if err := func() error {
			ownPeer := e
			return g.peers[i].eventHandler(tCtx,
				func(sctx context.Context, p *peer, sub event.Subscription) {
					var (
						err           error
						evt           interface{}
						nSent         = 0
						nRcv          = 0
						depPeer       = ownPeer - 1
						lastPeer      = peersAmount - 1
						stopCondition = fmt.Sprintf("%s_%s_msg%d", name, g.peers[peerIndexes[lastPeer]].name, msgAmount-1)
					)

					if depPeer == -1 {
						depPeer = lastPeer
					}

					defer func() {
						// Wait for last message to be synced among peers
						time.Sleep(time.Duration(peersAmount) * time.Second)

						sub.Close()

						errsMutex.Lock()
						errs[ownPeer] = err
						errsMutex.Unlock()

						wgEnd.Done()
					}()

					wgStart.Done()
					wgStart.Wait()
					time.Sleep(300 * time.Millisecond) // Wait for first sync

					if msgAmount == 0 {
						// Wait for timeout
						for {
							select {
							case <-sctx.Done():
								return
							case <-sub.Out():
								// Drain channel
							}
						}
					} else if ownPeer == 0 {
						time.Sleep(100 * time.Millisecond)
						p.sendMessage(sctx, fmt.Sprintf("%s_%s_msg%d", name, p.name, nSent))
						nSent++
					}

					for {
						waitCondition := fmt.Sprintf("%s_%s_msg%d", name, g.peers[peerIndexes[depPeer]].name, nRcv)
						select {
						case <-sctx.Done():
							err = sctx.Err()
							return
						case evt = <-sub.Out():
							erp, ok := evt.(stores.EventReplicateProgress)
							if outOfOrder || len(peerIndexes) == 1 {
								if nSent >= msgAmount {
									return
								}
								p.sendMessage(sctx, fmt.Sprintf("%s_%s_msg%d", name, p.name, nSent))
								nSent++
							} else if ok {
								msg := getEntryValue(erp.Entry)
								if msg == stopCondition {
									return
								} else if msg == waitCondition {
									nRcv++
									// time.Sleep(100 * time.Millisecond) // Avoid sending multiple message with same parent
									p.sendMessage(sctx, fmt.Sprintf("%s_%s_msg%d", name, p.name, nSent))
									// time.Sleep(50 * time.Millisecond) // Avoid sending multiple message with same parent
									nSent++
									if ownPeer == lastPeer && nSent >= msgAmount {
										return
									}
								}
							}
						}
					}
				})
		}(); err != nil {
			return err
		}
	}

	wgEnd.Wait()

	failed := false
	wrapped := fmt.Errorf("%s failed:", name)
	for i, err := range errs {
		if err != nil {
			failed = true
		}
		wrapped = fmt.Errorf("%s [%s: %v]", wrapped.Error(), g.peers[peerIndexes[i]].name, err)
	}
	if failed {
		return wrapped
	}

	return nil
}

func (g *group) close() {
	g.mocknet.Close()
	for _, peer := range g.peers {
		peer.close()
	}
}
