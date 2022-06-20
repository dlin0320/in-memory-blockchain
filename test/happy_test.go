package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/dlin0320/in-memory-blockchain/client"
	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
	"google.golang.org/grpc"
)

const (
	HappyRequests = 10
)

var bc *client.BlockchainClient
var conn *grpc.ClientConn
var ctx context.Context
var scheduler common.Scheduler

func setup() {
	bc, conn = client.Dial()
	ctx = context.Background()
	scheduler = common.NewScheduler(false, true)
}

func teardown() {
	conn.Close()
}

func getBlock() {
	param := pb.QueryParams{}
	list, err := bc.GetBlock(ctx, &param)
	if err != nil {
		log.Printf("error is: %v\n", err)
	} else {
		log.Printf("block is: %v\n", list)
	}
}

func happyTasks(c *client.BlockchainClient, ctx context.Context) []*common.Task {
	var task_list []*common.Task
	for i := 0; i < HappyRequests; i++ {
		t := common.NewTask(float64(i), func() {
			c.CreateTransaction(ctx, client.GenPayload())
		})
		task_list = append(task_list, t)
	}

	return task_list
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestHappy(t *testing.T) {
	log.Printf("running %v", t.Name())
	tasks := happyTasks(bc, ctx)
	for _, t := range tasks {
		scheduler.Schedule(t)
	}
	scheduler.Wait()
}
