package test

import (
	"context"

	"sort"
	"testing"
	"time"

	"golang.org/x/exp/rand"

	"github.com/dlin0320/in-memory-blockchain/client"
	"github.com/dlin0320/in-memory-blockchain/common"

	"gonum.org/v1/gonum/stat/distuv"
)

const (
	Min           = 0
	Max           = 20
	TotalRequests = 10
)

func uniform(c *client.BlockchainClient, ctx context.Context) []*common.Task {
	seed := time.Now().UnixNano()
	src := rand.NewSource(uint64(seed))
	u := distuv.Uniform{Min: Min, Max: Max, Src: src}
	var task_list []*common.Task
	for i := 0; i < TotalRequests; i++ {
		n := u.Rand()
		task := common.NewTask(n, func() {
			c.CreateTransaction(ctx, client.GenPayload())
		})
		task_list = append(task_list, task)
	}
	sort.Slice(task_list, func(i, j int) bool {
		return task_list[i].Time < task_list[j].Time
	})

	return task_list
}

func TestUniform(t *testing.T) {

	s := common.NewScheduler()
	client, conn := client.Dial()
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), Max*time.Second)
	defer cancel()

	tasks := uniform(client, ctx)
	for _, t := range tasks {
		s.Schedule(t)
	}
	s.Wait()
}
