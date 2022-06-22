package server

import (
	"log"
	"net"
	"sync"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
	"github.com/enriquebris/goconcurrentqueue"

	"google.golang.org/grpc"
)

var server *BlockchainServer

var bcChannel = make(chan BlockchainServer)

type Balance struct {
	final float64
	temp  float64
}

type BlockchainServer struct {
	pb.UnimplementedBlockchainServer
	tx_queue goconcurrentqueue.FIFO
	balance  sync.Map
	latest   *pb.Blk
	blocks   map[string]*pb.Blk
}

func newBlockchainServer() *BlockchainServer {
	return &BlockchainServer{blocks: map[string]*pb.Blk{}}
}

func Serve() {
	lis, err := net.Listen("tcp", common.ServerPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server = newBlockchainServer()
	bcChannel <- *server
	pb.RegisterBlockchainServer(grpcServer, server)
	log.Printf("server listening at %v\n", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func Mine() {
	<-bcChannel
	log.Println("start mining...")
	mine()
}
