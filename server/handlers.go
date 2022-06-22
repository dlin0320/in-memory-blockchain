package server

import (
	"context"
	"fmt"
	"log"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

func checkAddress(addr string) {
	balance, _ := server.balance.Load(addr)
	if balance == nil {
		log.Printf("observed new account: %v", addr)
		server.balance.Store(addr, &Balance{final: 100, temp: 0})
	}
}

func checkTx(in *pb.TxPayload) error {
	checkAddress(in.GetFrom())
	checkAddress(in.GetTo())
	sender, _ := server.balance.Load(in.GetFrom())
	balance := sender.(*Balance)
	if balance.final+balance.temp < in.GetValue() || in.GetFrom() == in.GetTo() {
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

func findTx(h string, starting_block string, search_range int) *pb.Tx {
	curr_block := server.blocks[starting_block]

	for i := 0; i < search_range; i++ {
		tx := findInBlock(curr_block, h)
		if tx != nil {
			return tx
		}
		parent := curr_block.GetHeader().GetParentHash()
		curr_block = server.blocks[parent]
	}

	return nil
}

func getBalance(addr string) (*Balance, error) {
	b, _ := server.balance.Load(addr)
	if b == nil {
		return nil, fmt.Errorf("account not found with address: %v", addr)
	}
	return b.(*Balance), nil
}

func (s *BlockchainServer) CreateTransaction(ctx context.Context, in *pb.TxPayload) (*pb.Tx, error) {
	log.Printf("received new tx: %v", in)

	// create new transaction and push to mempool
	var err error
	err = checkTx(in)
	if err != nil {
		return nil, err
	}
	tx := newTx(in)
	err = s.tx_queue.Enqueue(tx)
	if err != nil {
		return nil, err
	}

	// update temp balance
	sender_balance, _ := getBalance(tx.GetFrom())
	receiver_balance, _ := getBalance(tx.GetTo())
	sender_balance.temp -= tx.Value
	receiver_balance.temp += tx.Value

	log.Printf("new tx created: %v", tx)
	return tx, nil
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

		findTx(tx_hash, starting_block, search_range)
	}

	return tx_list, nil
}

func (s *BlockchainServer) GetBlock(ctx context.Context, in *pb.QueryParams) (*pb.BlkList, error) {
	log.Printf("getting latest block")
	blk_list := pb.BlkList{Blocks: []*pb.Blk{s.latest}}

	return &blk_list, nil
}

func (s *BlockchainServer) GetBalance(ctx context.Context, in *pb.QueryParams) (*pb.Balance, error) {
	log.Printf("getting balance for: %v", in.GetAddress())
	b, err := getBalance(in.GetAddress())
	if b == nil {
		return nil, err
	}
	balance := pb.Balance{Address: in.GetAddress(), Balance: b.final}
	return &balance, nil
}
