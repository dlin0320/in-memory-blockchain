package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

type BlockchainClient struct {
	pb.BlockchainClient
}

func newBlockchainClient(conn *grpc.ClientConn) *BlockchainClient {
	return &BlockchainClient{pb.NewBlockchainClient(conn)}
}

func Dial() (*BlockchainClient, *grpc.ClientConn) {

	conn, err := grpc.Dial(common.ServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := newBlockchainClient(conn)

	return client, conn
}
