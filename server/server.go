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

var server *BlockchainServer

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

func Serve() {
	lis, err := net.Listen("tcp", common.ServerPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server = newBlockchainServer()
	pb.RegisterBlockchainServer(grpcServer, server)
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func Mine() {
	fmt.Println("start mining...")
	server.mine()
}
