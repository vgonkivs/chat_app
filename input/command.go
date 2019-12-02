package input

import (
	"strings"

	"github.com/vgonkivs/ddd/pb"
)

func parseCommand(input string) pb.Message_Command {
	com := strings.TrimSuffix(input, " ")

	switch com {
	case "/list":
		return pb.Message_LIST
	case "/changeName":
		return pb.Message_CHANGE_NAME
	case "/info":
		return pb.Message_ACCOUNT
	case "/exit":
		return pb.Message_EXIT
	default:
		return pb.Message_NONE
	}
}
