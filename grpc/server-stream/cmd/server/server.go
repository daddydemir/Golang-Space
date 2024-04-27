package main

import (
	"errors"
	"log"
	"server-stream/proto/message"
	"server-stream/proto/service"
)

var videoChunks = [][]byte{
	[]byte("part-1"),
	[]byte("part-2"),
	[]byte("part-3"),
	[]byte("part-4"),
	[]byte("part-5"),
}

type Server struct {
	service.UnimplementedServerServiceServer
}

func (s Server) GetVideoStream(request *message.Request, server service.ServerService_GetVideoStreamServer) error {

	log.Printf("REQUEST: %v \n", request)

	if request.VideoId == "1337" {
		for _, c := range videoChunks {
			if err := server.Send(&message.Response{ChunkData: c}); err != nil {
				return err
			}
		}
	} else {
		return errors.New("video not found")
	}

	return nil
}
