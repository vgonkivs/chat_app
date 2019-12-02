package db

import (
	net "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

// DB is a simple interface to a database
type Repository interface {
	Store(p peer.ID, s net.Stream) error
	IsUserExist(p peer.ID) bool
}

type DB interface {
	GetUser(p peer.ID) (string, error)
	GetAllUsers() []string
	GetAllStreams() []net.Stream
	DeleteUser(p peer.ID) error
}
