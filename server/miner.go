package server

import (
	"fmt"
	"log"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

const (
	BlockTime = 10
)

func (s *BlockchainServer) updateAccountBalance(transactions []*pb.Tx) {
	for _, tx := range transactions {
		s.address_balance.final[common.StringToAddress(tx.GetFrom())] -= tx.GetValue()
		s.address_balance.final[common.StringToAddress(tx.GetTo())] += tx.GetValue()
	}
}

func (s *BlockchainServer) getTransactions() ([]*pb.Tx, error) {
	len := s.tx_queue.GetLen()
	tx_list := []*pb.Tx{}
	for i := 0; i < len; i++ {
		v, err := s.tx_queue.Dequeue()
		if err != nil {
			fmt.Println(err)
			return tx_list, err
		}
		tx, _ := v.(*pb.Tx)
		tx_list = append(tx_list, tx)
	}
	return tx_list, nil
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
		tx_hashes := [][]byte{}
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
		hash := common.GetHash(&block)
		block.Header.Hash = hash[:]
		log.Printf("new block created: %x", block.Header.Hash)

		// updating latest block and adding block to chain
		s.latest = &block
		s.blocks[string(block.Header.Hash)] = &block

		// update balances after block is added
		s.updateAccountBalance(transactions)
		log.Printf("new block mined: %x", block.Header.Hash)
	})

	scheduler.Schedule(task)
	scheduler.Wait()
}
