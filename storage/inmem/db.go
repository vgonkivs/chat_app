package inmem

import (
	"errors"
	"fmt"
	"sync"

	net "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	database "github.com/vgonkivs/ddd/db"
)

var errUnknownUser = errors.New("Unknown User")

type User struct {
	name   string
	stream net.Stream
}

// GetName ...
func (u *User) GetName() string {
	return u.name
}

// GetStream ...
func (u *User) GetStream() net.Stream {
	return u.stream
}

// NewUser creates new User
func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) String() string {
	return fmt.Sprintf("%v", u.name)
}

type db struct {
	storage  map[peer.ID]*User
	mu       *sync.RWMutex
	selfInfo *User
}

// NewInmemDB returns instane of local db
func NewInmemDB(p peer.ID) database.Repository {
	return &db{storage: make(map[peer.ID]*User), mu: &sync.RWMutex{}, selfInfo: &User{p.Pretty(), nil}}
}

func (db *db) IsUserExist(p peer.ID) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	_, ok := db.storage[p]
	return ok
}

// Store adds User in db
func (db *db) Store(p peer.ID, s net.Stream) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if p.Pretty() == db.selfInfo.name {
		return errors.New("Not able to add selInfo into Db")
	}
	if s == nil {
		return errors.New("Could not add the empty stream")
	}
	if _, ok := db.storage[p]; !ok {
		db.storage[p] = &User{name: p.Pretty(), stream: s}
		return nil
	}
	return errors.New("Already connected")
}

// DeleteUser ... deletes User from db
func (db *db) DeleteUser(p peer.ID) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, ok := db.storage[p]; ok {
		delete(db.storage, p)
		return nil
	}
	return errUnknownUser
}

// GetUser ...
func (db *db) GetUser(p peer.ID) (string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if p.Pretty() == db.selfInfo.name {
		return db.selfInfo.String(), nil
	}

	if user, ok := db.storage[p]; ok {
		return user.String(), nil
	}
	return "", errUnknownUser
}

// GetAllStreams ...
func (db *db) GetAllStreams() []net.Stream {
	db.mu.RLock()
	defer db.mu.RUnlock()
	arr := make([]net.Stream, 0)
	for _, value := range db.storage {
		arr = append(arr, value.stream)
	}
	return arr
}

// GetAllUsers ...
func (db *db) GetAllUsers() []string {
	db.mu.RLock()
	defer db.mu.RUnlock()
	list := make([]string, 0)
	for key := range db.storage {
		list = append(list, key.Pretty())
	}
	return list
}
