package server

import (
	"context"
	"log"
	"sync"

	"github.com/enriquebris/goconcurrentqueue"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

type AddressWithBalance struct {
	addr_bal_map sync.Map
}

type TransactionServer struct {
	pb.UnimplementedTransactionServer
	tx_queue        *goconcurrentqueue.FIFO
	tx_list         *pb.TxList
	address_balance AddressWithBalance
}

func (s *TransactionServer) checkAddress(add string) {
	s.address_balance.addr_bal_map.LoadOrStore(add, 100)
}

func NewTransactionServer() *TransactionServer {
	return &TransactionServer{tx_list: &pb.TxList{}}
}

func NewTx(p *pb.TxPayload) *pb.Tx {
	hash := common.GetHash(p)
	return &pb.Tx{Hash: hash[:], From: p.GetFrom(), To: p.GetTo(), Value: p.GetValue()}
}

func (s *TransactionServer) CreateTx(ctx context.Context, in *pb.TxPayload) (*pb.Tx, error) {
	log.Printf("received new tx: %v", in.String())
	go s.checkAddress(in.GetFrom())
	go s.checkAddress(in.GetTo())
	p := &pb.TxPayload{From: in.GetFrom(), To: in.GetTo(), Value: in.GetValue()}
	newTx := NewTx(p)
	log.Printf("new tx created: %s", newTx)
	s.tx_queue.Enqueue(newTx)

	return newTx, nil
}

func (s *TransactionServer) GetTx(ctx context.Context, in *pb.TxParams) (*pb.TxList, error) {
	return s.tx_list, nil
}
