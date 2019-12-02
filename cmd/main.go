package main

import (
	"context"

	multiaddr "github.com/multiformats/go-multiaddr"
	"github.com/vgonkivs/ddd"
)

func main() {
	network.NewNode(New()).RunApplication()
}

// New for test puproses
func New() (ctx context.Context, cancel context.CancelFunc, addr multiaddr.Multiaddr) {
	ctx, cancel = context.WithCancel(context.Background())
	a := "/ip4/127.0.0.1/tcp/0"
	addr, _ = multiaddr.NewMultiaddr(a)
	return ctx, cancel, addr
}
