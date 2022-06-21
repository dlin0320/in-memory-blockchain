package server

import (
	"fmt"
	"log"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

const (
	BlockTime = 10
	BlockSize = 500
)

func (s *BlockchainServer) updateAccountBalance(transactions []*pb.Tx) {
	for _, tx := range transactions {
		val := tx.GetValue()
		sender, _ := s.balance.Load(tx.GetFrom())
		receiver, _ := s.balance.Load(tx.GetTo())
		sender_balance := sender.(Balance)
		receiver_balance := receiver.(Balance)
		sender_balance.final -= val
		sender_balance.temp += val
		receiver_balance.final += val
		receiver_balance.temp -= val
		s.balance.Store(tx.GetFrom(), sender_balance)
		s.balance.Store(tx.GetTo(), receiver_balance)
	}
}

func (s *BlockchainServer) getTransactions() ([]*pb.Tx, error) {
	len := s.tx_queue.GetLen()
	if len > BlockSize {
		len = BlockSize
	}
	transactions := []*pb.Tx{}
	for i := 0; i < len; i++ {
		v, err := s.tx_queue.Dequeue()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tx, _ := v.(*pb.Tx)
		transactions = append(transactions, tx)
	}

	return transactions, nil
}

func (s *BlockchainServer) mine() {
	index := 0
	scheduler := common.NewScheduler(true, false)

	task := common.NewTask(BlockTime, func() {
		log.Printf("mining block number %v", index)
		index++

		// preparing data for the block
		transactions, err := s.getTransactions()
		if err != nil {
			log.Fatalf("failed to get transactions: %v", err)
		}
		tx_list := pb.TxList{Transactions: transactions}
		tx_hashes := []string{}
		for _, tx := range transactions {
			tx_hashes = append(tx_hashes, tx.Hash)
		}
		var header pb.Header
		if s.latest == nil {
			header = pb.Header{Height: int32(len(transactions)), TxHashes: tx_hashes}
		} else {
			header = pb.Header{ParentHash: s.latest.Header.Hash, Height: int32(len(transactions)), TxHashes: tx_hashes}
		}

		// generating the block's hash
		block := pb.Blk{TxList: &tx_list, Header: &header}
		block.Header.Hash = common.GetHash(&block)
		log.Printf("new block created: %v", block.Header.Hash)

		// updating latest block and adding block to chain
		s.latest = &block
		s.blocks[block.Header.Hash] = &block

		// update balances after block is added
		s.updateAccountBalance(transactions)
		log.Printf("new block mined: %v", block.Header.Hash)
	})

	scheduler.Schedule(task)
	scheduler.Wait()
}
