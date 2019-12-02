package endpoint

import (
	"fmt"
	"sync"

	"github.com/vgonkivs/ddd/serializer"

	net "github.com/libp2p/go-libp2p-core/network"
	msgio "github.com/libp2p/go-msgio"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/vgonkivs/ddd/pb"

	"github.com/vgonkivs/ddd/db"
)

type handlerFunc func() error

// Endpoints ...
type Endpoints interface {
	Do(mess *pb.Message) handlerFunc
}

type handler struct {
	db     db.DB
	selfID peer.ID
	mu     *sync.Mutex
}

// NewMux creates an instance of a router
func NewMux(db db.DB, p peer.ID) Endpoints {
	return &handler{db: db, selfID: p, mu: &sync.Mutex{}}
}

func (h *handler) list() string {
	return wrapList(h.db.GetAllUsers())
}

func (h *handler) accountInfo() (string, error) {
	return h.db.GetUser(h.selfID)
}

func (h *handler) deleteUser(pID string) error {
	peerID, err := peer.IDB58Decode(pID)
	if err != nil {
		return err
	}

	return h.db.DeleteUser(peerID)
}

func (h *handler) getStreams() []net.Stream {
	return h.db.GetAllStreams()
}

func (h *handler) writeMsgToStream(mess *pb.Message) error {
	streams := h.getStreams()
	mess.Command = pb.Message_COMMON

	wg := &sync.WaitGroup{}
	wg.Add(len(streams))
	buff := serializer.Marshal(mess)

	for _, stream := range streams {
		go func(w net.Stream) {
			defer wg.Done()
			err := msgio.NewWriter(w).WriteMsg((buff[:len(buff)]))
			if err != nil {
				h.db.DeleteUser(w.Conn().RemotePeer())
				w.Close()
			}
		}(stream)
	}
	wg.Wait()
	return nil
}

func (h *handler) writToOutput(str string) error {
	_, err := fmt.Print(str + "\n")
	return err
}

func (h *handler) Do(mess *pb.Message) handlerFunc {
	h.mu.Lock()
	defer h.mu.Unlock()

	switch mess.Command {
	case pb.Message_LIST:
		return func() error { return h.writToOutput(h.list()) }
	case pb.Message_ACCOUNT:
		return func() error {
			if str, err := h.accountInfo(); err == nil {
				return h.writToOutput(str)
			}
			return nil
		}
	case pb.Message_CHANGE_NAME:
	case pb.Message_PING:
		// not implemented yet
		return nil
	case pb.Message_COMMON:
		return func() error { return h.writToOutput(string(mess.GetContent())) }
	case pb.Message_EXIT:
		return func() error { return h.deleteUser(mess.GetFrom()) }
	default:
		return func() error { return h.writeMsgToStream(mess) }
	}
	return nil
}

func wrapList(list []string) string {
	str := "[ "
	for _, key := range list {
		str += key + " "
	}
	str += "]"
	return str
}
