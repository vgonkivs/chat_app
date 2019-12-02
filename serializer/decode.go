package serializer

import (
	"log"

	"github.com/gogo/protobuf/proto"

	pb "github.com/vgonkivs/ddd/pb"
)

// Marshal marshals message into bytes
func Marshal(msg *pb.Message) []byte {
	buff, err := proto.Marshal(msg)
	if err != nil {
		log.Println(err)
		return nil
	}
	return buff
}
