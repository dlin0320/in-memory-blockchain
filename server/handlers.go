package server

import (
	"context"
	"fmt"
	"log"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

func (s *BlockchainServer) checkAddress(addr string) {
	balance, _ := s.balance.Load(addr)
	if balance == nil {
		log.Printf("observed new account: %v", addr)
		s.balance.Store(addr, Balance{final: 0, temp: 100})
	}
}

func (s *BlockchainServer) checkTx(in *pb.TxPayload) error {
	s.checkAddress(in.GetFrom())
	s.checkAddress(in.GetTo())
	sender, found := s.balance.Load(in.GetFrom())
	balance := sender.(Balance)
	if !found || balance.final+balance.temp < in.GetValue() || in.GetFrom() == in.GetTo() {
		return fmt.Errorf("invalid transaction: %v", in)
	}
	return nil
}

func newTx(p *pb.TxPayload) *pb.Tx {
	hash := common.GetHash(p)
	return &pb.Tx{Hash: hash, From: p.GetFrom(), To: p.GetTo(), Value: p.GetValue()}
}

func findInBlock(b *pb.Blk, h string) *pb.Tx {
	hashes := b.GetHeader().GetTxHashes()
	var transactions []*pb.Tx
	for _, hash := range hashes {
		if hash == h {
			transactions = b.GetTxList().GetTransactions()
			break
		}
	}
	for _, tx := range transactions {
		tx_hash := tx.GetHash()
		if tx_hash == h {
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
		curr_block = s.blocks[parent]
	}

	return nil
}

func (s *BlockchainServer) CreateTransaction(ctx context.Context, in *pb.TxPayload) (*pb.Tx, error) {
	log.Printf("received new tx: %v", in)

	// create new transaction and push to mempool
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

	// update temp balance
	sender, _ := s.balance.Load(tx.GetFrom())
	receiver, _ := s.balance.Load(tx.GetTo())
	sender_balance := sender.(Balance)
	receiver_balance := receiver.(Balance)
	sender_balance.temp -= tx.Value
	receiver_balance.temp += tx.Value

	log.Printf("new tx created: %v", tx)
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
			starting_block = s.latest.GetHeader().GetHash()
		}
		search_range := int(in.GetRange())
		if search_range == 0 || search_range > 10 {
			search_range = 3
		}

		s.findTx(tx_hash, starting_block, search_range)
	}

	return tx_list, nil
}
