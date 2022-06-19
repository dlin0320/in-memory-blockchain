package server

import (
	"fmt"
	"time"

	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
)

func (s *BlockchainServer) getTransactions() ([]*pb.Tx, error) {
	len := s.tx_queue.GetLen()
	tx_list := []*pb.Tx{}
	for i := 0; i < len; i++ {
		v, err := s.tx_queue.Dequeue()
		if err != nil {
			fmt.Println(err)
			return tx_list, err
		}
		tx, ok := v.(*pb.Tx)
		if ok {
			tx_list = append(tx_list, tx)
		}
	}
	return tx_list, nil
}

func (s *BlockchainServer) Mine() {
	scheduler := common.NewScheduler()
	scheduler.ScheduleRecursive(func(again func()) {
		time.Sleep(10 * time.Second)
		s.getTransactions()
		again()
	})
}
