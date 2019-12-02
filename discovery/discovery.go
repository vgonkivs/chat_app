package discovery

import (
	"context"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/host"

	"github.com/libp2p/go-libp2p/p2p/discovery"
)

// DiscoveryService is a libp2p discovery service
type DiscoveryService interface {
	discovery.Service
}

// NewDiscoveryService creates a simple service to discover nodes in a local network
func NewDiscoveryService(host host.Host) DiscoveryService {
	discoverer, err := discovery.NewMdnsService(context.Background(), host, time.Second, "meetup") //node.config.RendezvousString) // TODO - change when rework config
	if err != nil {
		log.Println("Could not start discovery service: " + err.Error())
		return nil
	}

	return discoverer
}
