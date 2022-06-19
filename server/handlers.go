package server

import (
	"context"
	"log"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

func (s *BlockchainServer) checkAddress(addr string) {
	s.address_balance.LoadOrStore(addr, 100)
}

func newTx(p *pb.TxPayload) *pb.Tx {
	hash := common.GetHash(p)
	return &pb.Tx{Hash: hash[:], From: p.GetFrom(), To: p.GetTo(), Value: p.GetValue()}
}

func findInBlock(b *pb.Blk, h string) *pb.Tx {
	hashes := b.GetHeader().GetTxHashes()
	for _, hash := range hashes {
		if string(hash[:]) == h {
			tx := b.GetTransactions()[string(hash[:])]
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

	go s.checkAddress(in.GetFrom())
	go s.checkAddress(in.GetTo())
	newTx := newTx(in)
	s.tx_queue.Enqueue(newTx)

	log.Printf("new tx created: %x", newTx.Hash)
	return newTx, nil
}

func (s *BlockchainServer) GetBlock(ctx context.Context, in *pb.QueryParams) (*pb.BlkList, error) {
	blk_list := pb.BlkList{Blocks: []*pb.Blk{s.latest}}

	return &blk_list, nil
}

func (s *BlockchainServer) GetTransactions(ctx context.Context, in *pb.QueryParams) (*pb.TxList, error) {
	var tx_list *pb.TxList
	tx_hash := in.GetTxHash()

	if tx_hash == "" {
		transactions := s.latest.GetTransactions()
		for _, tx := range transactions {
			tx_list.Transactions = append(tx_list.Transactions, tx)
		}
	} else {
		starting_block := in.GetBlkHash()
		if starting_block == "" {
			starting_block = string(s.latest.GetHeader().GetHash()[:])
		}
		search_range := int(in.GetRange())
		if search_range == 0 {
			search_range = 3
		}

		s.findTx(tx_hash, starting_block, search_range)
	}

	return tx_list, nil
}
