package client

import (
	"context"
	"log"
	"math"
	"time"

	"golang.org/x/exp/rand"

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

func GenPayload() *pb.TxPayload {
	from := common.RandomHash(common.AddressLength)
	to := common.RandomHash(common.AddressLength)
	value := math.Round(rand.Float64()*100) / 100
	return &pb.TxPayload{From: from, To: to, Value: float32(value)}
}

func (c *BlockchainClient) CreateTransactions(payloads []*pb.TxPayload) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, v := range payloads {
		log.Printf("creating tx %v", v)
		tx, err := c.CreateTransaction(ctx, v)
		if err != nil {
			log.Printf("error creating tx: %v", err)
		} else {
			log.Printf("created tx: %v", tx)
		}
	}
}

func Dial() (*BlockchainClient, *grpc.ClientConn) {

	conn, err := grpc.Dial(common.ServerPort, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := newBlockchainClient(conn)

	return client, conn
}
