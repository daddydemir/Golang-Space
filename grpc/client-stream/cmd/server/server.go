package main

import (
	"client-stream/proto/message"
	"client-stream/proto/service"
	"io"
	"log"
)

type Server struct {
	service.UnimplementedClientServiceServer
}

func (s Server) UploadFile(server service.ClientService_UploadFileServer) error {

	var file []byte
	for {
		chunk, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error uploading file err: %v", err)
		}
		log.Printf("video part : %v", string(chunk.ChunkData))
		file = append(file, chunk.ChunkData...)
	}

	err := server.SendAndClose(&message.Response{Message: "file upload success"})
	if err != nil {
		log.Printf("Error uploading file err: %v", err)
	}

	log.Printf("REQUEST: %v", string(file))

	return nil
}
