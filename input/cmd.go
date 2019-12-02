package input

import (
	"bufio"
	"os"
	"strings"

	"github.com/vgonkivs/ddd/pb"
)

// Service ...
type Service interface {
	SetHandler(func(*pb.Message))
	EnableInput()
}

type inputService struct {
	handler func(*pb.Message)
}

// NewInputService creates a cli service
func NewInputService(handler func(*pb.Message)) Service {
	return &inputService{handler: handler}
}

// EnableInput ...
func (i *inputService) EnableInput() {
	r := bufio.NewReader(os.Stdin)
	for {
		str, err := r.ReadString('\n')
		if err != nil || str == "\n" {
		}
		comm := parseCommand(strings.Replace(str, "\n", "", -1))
		i.handler(prepareMessage(comm, strings.Replace(strings.TrimPrefix(str, " "), "\n", "", -1)))
	}
}

func prepareMessage(comm pb.Message_Command, content string) *pb.Message {
	return &pb.Message{Command: comm, Content: []byte(content)}
}

func (i *inputService) SetHandler(f func(*pb.Message)) {
	i.handler = f
}
