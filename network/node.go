package network

import (
	"context"
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
	"go-twing-sdk/common"
	"sync"
)

var topicNameFlag = flag.String("topicName", "twingdev", "twing development")

type Node struct {
	Cfg        *common.NetworkConfiguration
	ID         peer.ID
	ctx        context.Context
	cancelFunc context.CancelFunc
	Host       host.Host
	Streams    []*Stream
	conn       network.Conn
	pubsub     *pubsub.PubSub
	Topic      *pubsub.Topic
	Sub        *pubsub.Subscription
	gossip     *pubsub.GossipSubRouter
}

type Stream struct {
	muxed   network.MuxedStream
	stream  network.Stream
	handler network.StreamHandler
}

func NewNode() *Node {
	ctx := context.Background()
	//hostname, _ := os.Hostname()
	//common.NewNetworkConfiguration("twing-net", "1.0", hostname, "55435", "")
	ncfg := common.GetNetworkConfiguration()
	spew.Dump(ncfg)
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		common.ExceptionManager.ThrowFatalError(err.Error())
	}

	go discoverPeers(ctx, h)

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		common.ExceptionManager.ThrowFatalError(err.Error())
	}
	topic, err := ps.Join(*topicNameFlag)
	if err != nil {
		common.ExceptionManager.ThrowFatalError(err.Error())
	}

	sub, err := topic.Subscribe()
	if err != nil {
		common.ExceptionManager.ThrowFatalError(err.Error())
	}

	return &Node{
		Cfg:     ncfg,
		ID:      h.ID(),
		Host:    h,
		ctx:     ctx,
		Streams: make([]*Stream, 0),
		pubsub:  ps,
		Topic:   topic,
		Sub:     sub,
	}

}

func initDHT(ctx context.Context, h host.Host) *dht.IpfsDHT {
	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	kademliaDHT, err := dht.New(ctx, h)
	if err != nil {
		panic(err)
	}
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, peerAddr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Connect(ctx, *peerinfo); err != nil {
				fmt.Println("Bootstrap warning:", err)
			}
		}()
	}
	wg.Wait()

	return kademliaDHT
}

func discoverPeers(ctx context.Context, h host.Host) {
	kademliaDHT := initDHT(ctx, h)
	routingDiscovery := drouting.NewRoutingDiscovery(kademliaDHT)
	dutil.Advertise(ctx, routingDiscovery, *topicNameFlag)

	// Look for others who have announced and attempt to connect to them
	anyConnected := false
	for !anyConnected {
		fmt.Println("Searching for peers...")
		peerChan, err := routingDiscovery.FindPeers(ctx, *topicNameFlag)
		if err != nil {
			panic(err)
		}
		for peer := range peerChan {
			if peer.ID == h.ID() {
				continue // No self connection
			}
			err := h.Connect(ctx, peer)
			if err != nil {
				fmt.Printf("Failed connecting to %s, error: %s\n", peer.ID, err)
			} else {
				fmt.Println("Connected to:", peer.ID)
				anyConnected = true
			}
		}
	}
	fmt.Println("Peer discovery complete")
}
