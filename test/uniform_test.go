package test

import (
	"context"
	"log"

	"sort"
	"testing"
	"time"

	"golang.org/x/exp/rand"

	"github.com/dlin0320/in-memory-blockchain/client"
	"github.com/dlin0320/in-memory-blockchain/common"

	"gonum.org/v1/gonum/stat/distuv"
)

const (
	Min             = 0
	Max             = 20
	UniformRequests = 20
)

func uniformTasks(c *client.BlockchainClient, ctx context.Context) []*common.Task {
	seed := time.Now().UnixNano()
	src := rand.NewSource(uint64(seed))
	u := distuv.Uniform{Min: Min, Max: Max, Src: src}
	var task_list []*common.Task
	for i := 0; i < UniformRequests; i++ {
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
	log.Printf("running %v", t.Name())
	tasks := uniformTasks(bc, ctx)
	for _, t := range tasks {
		scheduler.Schedule(t)
	}
	scheduler.Wait()
}
