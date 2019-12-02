package serializer

import (
	"log"

	"github.com/gogo/protobuf/proto"
	pb "github.com/vgonkivs/ddd/pb"
)

// Unmarshal encodes data into pb.Message
func Unmarshal(buff []byte) *pb.Message {
	var mess pb.Message
	err := proto.Unmarshal(buff, &mess)
	if err != nil {
		log.Println(err)
	}

	return &mess
}
