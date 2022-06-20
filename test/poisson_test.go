package test

import (
	"context"
	"log"
	"sort"
	"testing"
	"time"

	"github.com/dlin0320/in-memory-blockchain/client"
	"github.com/dlin0320/in-memory-blockchain/common"
	"golang.org/x/exp/rand"
	dist "gonum.org/v1/gonum/stat/distuv"
)

const (
	Lambda          float64 = 50
	PoissonRequests int     = 100
)

func poissonTasks(c *client.BlockchainClient, ctx context.Context) []*common.Task {
	seed := time.Now().UnixNano()
	src := rand.NewSource(uint64(seed))
	p := dist.Poisson{Lambda: Lambda, Src: src}
	var task_list []*common.Task
	for i := 0; i < PoissonRequests; i++ {
		n := p.Rand()
		task := common.NewTask(n, func() {
			c.CreateTransaction(ctx, genPayload())
		})
		task_list = append(task_list, task)
	}

	sort.Slice(task_list, func(i, j int) bool {
		return task_list[i].Time < task_list[j].Time
	})
	return task_list
}

func TestPoisson(t *testing.T) {
	log.Printf("running %v", t.Name())
	tasks := poissonTasks(bc, ctx)
	for _, t := range tasks {
		scheduler.Schedule(t)
	}
	scheduler.Wait()
}
