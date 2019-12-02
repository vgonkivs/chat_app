package network

import (
	"context"
	"log"
	"os"
	"os/signal"

	protocol "github.com/libp2p/go-libp2p-core/protocol"

	"github.com/vgonkivs/ddd/db"
	"github.com/vgonkivs/ddd/endpoint"
	"github.com/vgonkivs/ddd/input"
	"github.com/vgonkivs/ddd/storage/inmem"

	net "github.com/libp2p/go-libp2p-core/network"

	"github.com/vgonkivs/ddd/discovery"
	"github.com/vgonkivs/ddd/pb"
	"github.com/vgonkivs/ddd/serializer"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	msgio "github.com/libp2p/go-msgio"
	tcp "github.com/libp2p/go-tcp-transport"

	"github.com/libp2p/go-libp2p-core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
)

// Network ...
type Network struct {
	ctx    context.Context
	cancel context.CancelFunc

	selfID peer.ID

	host host.Host

	discovery discovery.DiscoveryService
	db        db.Repository
	handlers  endpoint.Endpoints
	input     input.Service
}

// NewNode ...
func NewNode(ctx context.Context, cancel context.CancelFunc, addr multiaddr.Multiaddr) *Network {

	host, err := libp2p.New(ctx, libp2p.ListenAddrs(addr), libp2p.ChainOptions(libp2p.Transport(tcp.NewTCPTransport)))
	if err != nil {
		log.Println("Not ready to create a host ", err)
		cancel()
		return nil
	}
	database := inmem.NewInmemDB(host.ID())
	return &Network{
		ctx:       ctx,
		cancel:    cancel,
		selfID:    host.ID(),
		host:      host,
		discovery: discovery.NewDiscoveryService(host),
		db:        database,
		handlers:  endpoint.NewMux((database).(db.DB), host.ID()),
	}
}

func (n *Network) HandlePeerFound(pid peerstore.PeerInfo) {
	if n.db.IsUserExist(pid.ID) || n.connect(pid) != nil {
		return
	}
	n.dial(pid)
}

func (n *Network) RunApplication() {
	c := make(chan os.Signal)
	ch := make(chan struct{})
	defer func() {
		close(ch)
		close(c)
	}()

	n.input = input.NewInputService(n.handleInput)
	//n.input.SetHandler(n.handleInput)
	go n.input.EnableInput()

	n.host.SetStreamHandler(protocol.ID("/chat/1.0"), n.handleNewStream)
	n.discovery.RegisterNotifee(n)

	go func() {
		signal.Notify(c, os.Interrupt)
		select {
		case <-c:
			n.Close()
			ch <- struct{}{}

		}
	}()
	<-ch
	return
}

func (n *Network) Close() {
	n.db = nil
	n.host.Close()
	n.cancel()
}

func (n *Network) connect(pi peer.AddrInfo) error {
	return n.host.Connect(n.ctx, pi)
}

func (n *Network) handleNewStream(s net.Stream) {
	rc := msgio.NewReader(s)
	defer rc.Close()

	for {
		buff, err := rc.ReadMsg()
		if err != nil {
			if n.db == nil {
				break
			}
			n.handleInput(&pb.Message{Command: pb.Message_EXIT, From: s.Conn().RemotePeer().Pretty()})
			break
		}
		// Check with race detector  as handleInput + handleStream is handling on different routines
		go n.handleInput(serializer.Unmarshal(buff))
	}
}

func (n *Network) handleInput(mess *pb.Message) {
	f := n.handlers.Do(mess)
	if err := f(); err != nil {
		log.Println(err)
	}
}

func (n *Network) dial(pid peerstore.PeerInfo) {

	s, err := n.host.NewStream(context.Background(), pid.ID, protocol.ID("/chat/1.0"))
	if err != nil {
		return
	}
	_ = n.db.Store(pid.ID, s)

}
