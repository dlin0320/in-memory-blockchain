package server

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
	"github.com/enriquebris/goconcurrentqueue"

	"google.golang.org/grpc"
)

type BlockchainServer struct {
	pb.UnimplementedBlockchainServer
	tx_queue        goconcurrentqueue.FIFO
	address_balance sync.Map
	latest          *pb.Blk
	blocks          map[string]*pb.Blk
}

func newBlockchainServer() *BlockchainServer {
	return &BlockchainServer{}
}

func Serve() *BlockchainServer {
	lis, err := net.Listen("tcp", common.ServerPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	txServer := newBlockchainServer()
	pb.RegisterBlockchainServer(grpcServer, txServer)
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return txServer
}
