package server

import (
	"context"
	"fmt"
	"log"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

func (s *BlockchainServer) checkAddress(addr string) {
	_, exists := s.address_balance.current.LoadOrStore(addr, float32(100))
	if !exists {
		log.Printf("observed new account: %v", addr)
	}
}

func (s *BlockchainServer) checkTx(in *pb.TxPayload) error {
	s.checkAddress(in.GetFrom())
	s.checkAddress(in.GetTo())
	sender_balance, found := s.address_balance.current.Load(in.GetFrom())
	if !found || sender_balance.(float32) < in.GetValue() || in.GetFrom() == in.GetTo() {
		return fmt.Errorf("invalid transaction: %v", in)
	}
	return nil
}

func newTx(p *pb.TxPayload) *pb.Tx {
	hash := common.GetHash(p)
	return &pb.Tx{Hash: hash[:], From: p.GetFrom(), To: p.GetTo(), Value: p.GetValue()}
}

func findInBlock(b *pb.Blk, h string) *pb.Tx {
	hashes := b.GetHeader().GetTxHashes()
	var transactions []*pb.Tx
	for _, hash := range hashes {
		if string(hash[:]) == h {
			transactions = b.GetTxList().GetTransactions()
			break
		}
	}
	for _, tx := range transactions {
		tx_hash := tx.GetHash()
		if string(tx_hash[:]) == h {
			return tx
		}
	}
	return nil
}

func (s *BlockchainServer) findTx(h string, starting_block string, search_range int) *pb.Tx {
	curr_block := s.blocks[starting_block]

	for i := 0; i < search_range; i++ {
		tx := findInBlock(curr_block, h)
		if tx != nil {
			return tx
		}
		parent := curr_block.GetHeader().GetParentHash()
		curr_block = s.blocks[string(parent[:])]
	}

	return nil
}

func (s *BlockchainServer) CreateTransaction(ctx context.Context, in *pb.TxPayload) (*pb.Tx, error) {
	log.Printf("received new tx: %v", in.String())

	var err error
	err = s.checkTx(in)
	if err != nil {
		return nil, err
	}
	tx := newTx(in)
	err = s.tx_queue.Enqueue(tx)
	if err != nil {
		return nil, err
	}

	log.Printf("new tx created: %x", tx.GetHash())
	return tx, nil
}

func (s *BlockchainServer) GetBlock(ctx context.Context, in *pb.QueryParams) (*pb.BlkList, error) {
	log.Printf("getting latest block")
	blk_list := pb.BlkList{Blocks: []*pb.Blk{s.latest}}
	log.Printf("latest block is: %s", &blk_list)

	return &blk_list, nil
}

func (s *BlockchainServer) GetTransactions(ctx context.Context, in *pb.QueryParams) (*pb.TxList, error) {
	var tx_list *pb.TxList
	tx_hash := in.GetTxHash()

	if tx_hash == "" {
		transactions := s.latest.GetTxList().GetTransactions()
		tx_list.Transactions = transactions
	} else {
		starting_block := in.GetBlkHash()
		if starting_block == "" {
			starting_block = string(s.latest.GetHeader().GetHash()[:])
		}
		search_range := int(in.GetRange())
		if search_range == 0 || search_range > 10 {
			search_range = 3
		}

		s.findTx(tx_hash, starting_block, search_range)
	}

	return tx_list, nil
}
