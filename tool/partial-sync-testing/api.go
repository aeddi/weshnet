package main

import (
	"context"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"path/filepath"

	ipfs_cfg "github.com/ipfs/kubo/config"
	ipfs_core "github.com/ipfs/kubo/core"
	ipfs_mock "github.com/ipfs/kubo/core/mock"
	ipfs_p2p "github.com/ipfs/kubo/core/node/libp2p"
	ipfs_repo "github.com/ipfs/kubo/repo"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	p2p_ci "github.com/libp2p/go-libp2p/core/crypto"
	p2pnetwork "github.com/libp2p/go-libp2p/core/network"
	p2p_peer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	p2p_mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/pkg/errors"

	"berty.tech/weshnet/pkg/ipfsutil"
	ipfs_mobile "berty.tech/weshnet/pkg/ipfsutil/mobile"
	tinder "berty.tech/weshnet/pkg/tinder"
	ipfs_fsrepo "github.com/ipfs/kubo/repo/fsrepo"
)

func loadRepoFromPath(path string) ipfs_repo.Repo {
	dir, _ := filepath.Split(path)
	if _, err := ipfsutil.LoadPlugins(dir); err != nil {
		panic(errors.Wrap(err, "failed to load plugins"))
	}

	if !ipfs_fsrepo.IsInitialized(path) {
		c := ipfs_cfg.Config{}
		priv, _, err := p2p_ci.GenerateKeyPairWithReader(p2p_ci.RSA, 2048, crand.Reader)
		if err != nil {
			panic(fmt.Sprintf("failed to generate pair key: %v", err))
		}

		pid, err := p2p_peer.IDFromPublicKey(priv.GetPublic())
		if err != nil {
			panic(fmt.Sprintf("failed to get pid from pub key: %v", err))
		}

		privkeyb, err := p2p_ci.MarshalPrivateKey(priv)
		if err != nil {
			panic(fmt.Sprintf("failed to get raw priv key: %v", err))
		}

		c.Bootstrap = []string{}
		c.Addresses.Swarm = []string{"/ip6/::/tcp/0"}
		c.Identity.PeerID = pid.Pretty()
		c.Identity.PrivKey = base64.StdEncoding.EncodeToString(privkeyb)

		c.Mounts = ipfs_cfg.Mounts{
			IPFS: "/ipfs",
			IPNS: "/ipns",
		}

		c.Ipns = ipfs_cfg.Ipns{
			ResolveCacheSize: 128,
		}

		c.Reprovider = ipfs_cfg.Reprovider{
			Interval: "12h",
			Strategy: "all",
		}

		c.Datastore = ipfs_cfg.Datastore{
			StorageMax:         "10GB",
			StorageGCWatermark: 90, // 90%
			GCPeriod:           "1h",
			BloomFilterSize:    0,
			Spec: map[string]interface{}{
				"type": "mount",
				"mounts": []interface{}{
					map[string]interface{}{
						"mountpoint": "/blocks",
						"type":       "measure",
						"prefix":     "flatfs.datastore",
						"child": map[string]interface{}{
							"type":      "flatfs",
							"path":      "blocks",
							"sync":      true,
							"shardFunc": "/repo/flatfs/shard/v1/next-to-last/2",
						},
					},
					map[string]interface{}{
						"mountpoint": "/",
						"type":       "measure",
						"prefix":     "leveldb.datastore",
						"child": map[string]interface{}{
							"type":        "levelds",
							"path":        "datastore",
							"compression": "none",
						},
					},
				},
			},
		}

		if err := ipfs_fsrepo.Init(path, &c); err != nil {
			panic(fmt.Errorf("failed to init ipfs repo: %w", err))
		}
	}

	rep, err := ipfs_fsrepo.Open(path)
	if err != nil {
		panic(fmt.Errorf("failed to open ipfs repo: %w", err))
	}

	return rep
}

func testingCoreAPIUsingMockNet(ctx context.Context, repoPath string, opts *ipfsutil.TestingAPIOpts) ipfsutil.CoreAPIMock {
	repo := loadRepoFromPath(repoPath)
	mrepo := ipfs_mobile.NewRepoMobile("", repo)

	mnode, err := ipfsutil.NewIPFSMobile(ctx, mrepo, &ipfsutil.MobileOptions{
		HostOption:    ipfs_mock.MockHostOption(opts.Mocknet),
		RoutingOption: ipfs_p2p.NilRouterOption,
		ExtraOpts: map[string]bool{
			"pubsub": false,
		},
	})
	if err != nil {
		panic(err)
	}
	h := mnode.PeerHost()

	mockDriver := opts.DiscoveryServer.Client(h)

	stinder, err := tinder.NewService(h, opts.Logger, mockDriver)
	discAdaptater := tinder.NewDiscoveryAdaptater(opts.Logger, stinder)

	ps, err := pubsub.NewGossipSub(ctx, h, pubsub.WithDiscovery(discAdaptater))
	exapi, err := ipfsutil.NewExtendedCoreAPIFromNode(mnode.IpfsNode)

	psapi := ipfsutil.NewPubSubAPI(ctx, opts.Logger, ps)
	exapi = ipfsutil.InjectPubSubCoreAPIExtendedAdapter(exapi, psapi)
	ipfsutil.EnableConnLogger(ctx, opts.Logger, mnode.PeerHost())

	api := &coreAPIMock{
		coreapi: exapi,
		mocknet: opts.Mocknet,
		pubsub:  ps,
		node:    mnode.IpfsNode,
		tinder:  stinder,
	}

	return api
}

type coreAPIMock struct {
	coreapi ipfsutil.ExtendedCoreAPI
	pubsub  *pubsub.PubSub
	mocknet p2p_mocknet.Mocknet
	node    *ipfs_core.IpfsNode
	tinder  *tinder.Service
}

func (m *coreAPIMock) ConnMgr() ipfsutil.ConnMgr {
	return m.node.PeerHost.ConnManager()
}

func (m *coreAPIMock) NewStream(ctx context.Context, p p2p_peer.ID, pids ...protocol.ID) (p2pnetwork.Stream, error) {
	return m.node.PeerHost.NewStream(ctx, p, pids...)
}

func (m *coreAPIMock) SetStreamHandler(pid protocol.ID, handler p2pnetwork.StreamHandler) {
	m.node.PeerHost.SetStreamHandler(pid, handler)
}

func (m *coreAPIMock) RemoveStreamHandler(pid protocol.ID) {
	m.node.PeerHost.RemoveStreamHandler(pid)
}

func (m *coreAPIMock) API() ipfsutil.ExtendedCoreAPI {
	return m.coreapi
}

func (m *coreAPIMock) MockNetwork() p2p_mocknet.Mocknet {
	return m.mocknet
}

func (m *coreAPIMock) MockNode() *ipfs_core.IpfsNode {
	return m.node
}

func (m *coreAPIMock) PubSub() *pubsub.PubSub {
	return m.pubsub
}

func (m *coreAPIMock) Tinder() *tinder.Service {
	return m.tinder
}

func (m *coreAPIMock) Close() {
	m.node.Close()
}
