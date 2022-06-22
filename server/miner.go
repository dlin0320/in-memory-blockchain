package server

import (
	"log"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

const (
	BlockTime = 10
	BlockSize = 1000
)

func updateAccountBalance(transactions []*pb.Tx) {
	for _, tx := range transactions {
		val := tx.GetValue()
		sender_balance, _ := getBalance(tx.GetFrom())
		receiver_balance, _ := getBalance(tx.GetTo())
		sender_balance.final -= val
		sender_balance.temp += val
		receiver_balance.final += val
		receiver_balance.temp -= val
		server.balance.Store(tx.GetFrom(), sender_balance)
		server.balance.Store(tx.GetTo(), receiver_balance)
	}
}

func updateBalanceTemp(transactions []*pb.Tx) {
	for _, tx := range transactions {
		val := tx.GetValue()
		sender_balance, _ := getBalance(tx.GetFrom())
		receiver_balance, _ := getBalance(tx.GetTo())
		sender_balance.temp += val
		receiver_balance.temp -= val
		server.balance.Store(tx.GetFrom(), sender_balance)
		server.balance.Store(tx.GetTo(), receiver_balance)
	}
}

func getTransactions() ([]*pb.Tx, error) {
	len := server.tx_queue.GetLen()
	if len > BlockSize {
		len = BlockSize
	}
	transactions := []*pb.Tx{}
	for i := 0; i < len; i++ {
		v, err := server.tx_queue.Dequeue()
		if err != nil {
			return nil, err
		}
		tx, _ := v.(*pb.Tx)
		transactions = append(transactions, tx)
	}

	return transactions, nil
}

func checkTransactions(transactions []*pb.Tx) ([]*pb.Tx, []*pb.Tx) {
	sender_balances := map[string]float64{}
	validTransactions := []*pb.Tx{}
	invalidTransactions := []*pb.Tx{}
	for _, tx := range transactions {
		sender := tx.GetFrom()
		_, found := sender_balances[sender]
		if !found {
			b, _ := getBalance(sender)
			sender_balances[sender] = b.final
		}

		if sender_balances[sender]-tx.GetValue() >= 0 {
			sender_balances[sender] -= tx.GetValue()
			validTransactions = append(validTransactions, tx)
		} else {
			log.Printf("invalid transaction revoked: %v", tx)
			invalidTransactions = append(invalidTransactions, tx)
		}
	}

	return validTransactions, invalidTransactions
}

func mine() {
	index := 0
	scheduler := common.NewScheduler(true, false)

	task := common.NewTask(BlockTime, func() {
		log.Printf("mining block number %v", index)
		index++

		// preparing data for the block
		transactions, err := getTransactions()
		if err != nil {
			log.Fatalf("failed to get transactions: %v", err)
		}
		validTransactions, invalidTransactions := checkTransactions(transactions)
		tx_list := pb.TxList{Transactions: validTransactions}
		tx_hashes := []string{}
		for _, tx := range validTransactions {
			tx_hashes = append(tx_hashes, tx.GetHash())
		}
		header := pb.Header{Height: int32(len(validTransactions)), TxHashes: tx_hashes}
		if server.latest != nil {
			header.ParentHash = server.latest.GetHeader().GetHash()
		}

		// generating the block's hash
		block := pb.Blk{TxList: &tx_list, Header: &header}
		block.GetHeader().Hash = common.GetHash(&block)
		log.Printf("new block created: %v", &block)

		// updating latest block and adding block to chain
		server.latest = &block
		server.blocks[block.GetHeader().GetHash()] = &block

		// update balances after block is added
		updateAccountBalance(validTransactions)
		updateBalanceTemp(invalidTransactions)
		log.Printf("new block mined: %v", block.GetHeader().GetHash())
	})

	scheduler.Schedule(task)
	scheduler.Wait()
}
