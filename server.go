package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/dlin0320/in-memory-blockchain/proto"
	s "github.com/dlin0320/in-memory-blockchain/server"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	txServer := s.NewTransactionServer()

	pb.RegisterTransactionServer(grpcServer, txServer)
	fmt.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
