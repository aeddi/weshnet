package main

import (
	"context"
	"errors"
	"path/filepath"

	odb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/baseorbitdb"
	"berty.tech/go-orbit-db/iface"
	"berty.tech/go-orbit-db/stores"
	"berty.tech/go-orbit-db/stores/operation"
	"berty.tech/weshnet/pkg/ipfsutil"
	"berty.tech/weshnet/pkg/tinder"

	datastore "github.com/ipfs/go-datastore"
	ds_sync "github.com/ipfs/go-datastore/sync"
	"github.com/libp2p/go-libp2p/core/event"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"go.uber.org/zap"
)

type peer struct {
	name     string
	storeDir string
	ipfs     ipfsutil.CoreAPIMock
	ds       *ds_sync.MutexDatastore
	odb      iface.OrbitDB
	logger   *zap.Logger
	store    iface.Store
}

func createNewPeer(ctx context.Context, name, storeDir string, mn mocknet.Mocknet, msrv *tinder.MockDriverServer, logger *zap.Logger) *peer {
	storeDir = filepath.Join(storeDir, name)
	ipfsDir := filepath.Join(storeDir, "ipfs")
	odbDir := filepath.Join(storeDir, "orbitdb")
	ds := ds_sync.MutexWrap(datastore.NewMapDatastore())
	logger = logger.Named(name)

	ipfsopts := &ipfsutil.TestingAPIOpts{
		Logger:          logger,
		Mocknet:         mn,
		DiscoveryServer: msrv,
		Datastore:       ds,
	}
	ipfs := testingCoreAPIUsingMockNet(ctx, ipfsDir, ipfsopts)

	odbOptions := &baseorbitdb.NewOrbitDBOptions{
		Logger:    logger,
		Directory: &odbDir,
	}

	orbitdb, err := odb.NewOrbitDB(ctx, ipfs.API(), odbOptions)
	if err != nil {
		panic(err)
	}

	return &peer{
		name:     name,
		storeDir: storeDir,
		ipfs:     ipfs,
		ds:       ds,
		odb:      orbitdb,
		logger:   logger,
	}
}

func (p *peer) sendMessage(ctx context.Context, msg string) error {
	op := operation.NewOperation(nil, "ADD", []byte(msg))
	e, err := p.store.AddOperation(ctx, op, nil)
	if err != nil {
		return err
	} else if !e.IsValid() {
		return errors.New("entry is invalid")
	}

	return nil
}

func (p *peer) eventHandler(ctx context.Context, handler func(context.Context, *peer, event.Subscription)) error {
	sub, err := p.store.EventBus().Subscribe([]interface{}{
		new(stores.EventReady),
		new(stores.EventNewPeer),
		new(stores.EventLoad),
		new(stores.EventLoadProgress),
		new(stores.EventReplicate),
		new(stores.EventReplicateProgress),
		new(stores.EventReplicated),
		new(stores.EventWrite),
	})
	if err != nil {
		return err
	}

	go handler(ctx, p, sub)

	return nil
}

func (p *peer) close() {
	p.ds.Close()
	p.ipfs.Close()
	p.odb.Close()
}
