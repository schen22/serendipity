package chat

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
)

// chatServiceServer is implementation of ChatServiceServer proto interface
type chatServiceServer struct {
	msg chan string
}

// NewChatServiceServer creates Chat service object
func NewChatServiceServer() ChatServiceServer {
	return &chatServiceServer{msg: make(chan string, 1000)}
}

// Send sends message to the server
func (s *chatServiceServer) Send(ctx context.Context, message *wrappers.StringValue) (*empty.Empty, error) {
	if message != nil {
		log.Printf("Send requested: message=%v", *message)
		s.msg <- message.Value
	} else {
		log.Print("Send requested: message=<empty>")
	}

	return &empty.Empty{}, nil
}

// Subscribe is streaming method to get echo messages from the server
func (s *chatServiceServer) Subscribe(e *empty.Empty, stream ChatService_SubscribeServer) error {
	log.Print("Subscribe requested")
	for {
		m := <-s.msg
		n := Message{Text: fmt.Sprintf("I have received from you: %s. Thanks!", m)}
		if err := stream.Send(&n); err != nil {
			// put message back to the channel
			s.msg <- m
			log.Printf("Stream connection failed: %v", err)
			return nil
		}
		log.Printf("Message sent: %+v", n)
	}
}
