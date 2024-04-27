package main

import (
	"bidirectional-stream/proto/message"
	"bidirectional-stream/proto/service"
	"io"
	"log"
)

type Server struct {
	service.UnimplementedBidirectionalStreamServer
}

func (s Server) StartChat(server service.BidirectionalStream_StartChatServer) error {
	for {
		req, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Recv error: %v", err)
		}

		log.Printf("SERVER: %v\n", req.Message)

		resp := &message.Response{Message: "from server."}
		if err = server.Send(resp); err != nil {
			log.Printf("Send error: %v", err)
		}
	}
}
