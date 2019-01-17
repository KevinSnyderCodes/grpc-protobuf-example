package main

import (
	"log"
	"net"

	protoplane "github.com/kevinsnydercodes/protobuf-example/protobuf/compiled/go/plane"
	"github.com/kevinsnydercodes/protobuf-example/server/plane"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func run() error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	protoplane.RegisterPlaneServer(server, &plane.Server{})

	return server.Serve(listener)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
