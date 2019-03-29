package main

import (
	"context"
	"log"
	"net"

	pinger "github.com/Asuforce/grpc-sample/pinger/lib"
	"google.golang.org/grpc"
)

func main() {
	listner, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}

	grpcSrv := grpc.NewServer()
	pinger.RegisterPingerServer(grpcSrv, &server{})
	log.Printf("Pinger server is running!")
	grpcSrv.Serve(listner)
}

type server struct{}

func (s *server) Ping(ctx context.Context, req *pinger.Empty) (*pinger.Pong, error) {
	pong := &pinger.Pong{
		Text: "pong",
	}
	return pong, nil
}
