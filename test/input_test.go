package test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dlin0320/in-memory-blockchain/client"
	"github.com/dlin0320/in-memory-blockchain/common"
	pb "github.com/dlin0320/in-memory-blockchain/proto"
	"google.golang.org/grpc"
)

var bc pb.BlockchainClient
var conn *grpc.ClientConn
var ctx context.Context
var cancel context.CancelFunc
var scheduler common.Scheduler
var addresses []string

func setup() {
	bc, conn = client.Dial()
	ctx, cancel = context.WithCancel(context.Background())
	scheduler = common.NewScheduler(false, true)
}

func teardown() {
	conn.Close()
	cancel()
}

func genAddresses(n int) []string {
	var generated []string
	for i := 0; i < n; i++ {
		hash := common.RandomHash(common.AddressLength)
		addresses = append(addresses, hash)
		generated = append(generated, hash)
	}
	return generated
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestHappy(t *testing.T) {
	log.Printf("running %v", t.Name())
	count := 10
	senders := genAddresses(count)
	receivers := genAddresses(count)
	var tasks []*common.Task
	for i := 0; i < count; i++ {
		j := i
		task := common.NewTask(float64(0), func() {
			bc.CreateTransaction(ctx, &pb.TxPayload{From: senders[j], To: receivers[j], Value: 10})
		})
		tasks = append(tasks, task)
	}

	for _, t := range tasks {
		scheduler.Schedule(t)
	}
	scheduler.Wait()
	time.Sleep(10 * time.Second)
	for i := 0; i < count; i++ {
		sender, _ := bc.GetBalance(ctx, &pb.QueryParams{Address: senders[i]})
		receiver, _ := bc.GetBalance(ctx, &pb.QueryParams{Address: receivers[i]})
		assert.Equal(t, float64(90), sender.Balance)
		assert.Equal(t, float64(110), receiver.Balance)
	}
}

func TestConcurrency(t *testing.T) {
	log.Printf("running %v", t.Name())
	times := 5
	count := 5
	senders := genAddresses(count)
	receivers := genAddresses(count)
	var tasks []*common.Task
	for i := 0; i < times; i++ {
		_i := i
		for j := 0; j < count; j++ {
			_j := j
			task := common.NewTask(float64(_i), func() {
				bc.CreateTransaction(ctx, &pb.TxPayload{From: senders[_j], To: receivers[_j], Value: 20})
			})
			tasks = append(tasks, task)
		}
	}

	for _, t := range tasks {
		scheduler.Schedule(t)
	}
	scheduler.Wait()

	time.Sleep(10 * time.Second)
	for i := 0; i < count; i++ {
		sender, _ := bc.GetBalance(ctx, &pb.QueryParams{Address: senders[i]})
		receiver, _ := bc.GetBalance(ctx, &pb.QueryParams{Address: receivers[i]})
		assert.Equal(t, float64(0), sender.Balance)
		assert.Equal(t, float64(200), receiver.Balance)
	}
}

func TestInvalidTx(t *testing.T) {
	log.Printf("running %v", t.Name())
	sender := common.RandomHash(common.AddressLength)
	receiver := common.RandomHash(common.AddressLength)
	task := common.NewTask(0, func() {
		bc.CreateTransaction(ctx, &pb.TxPayload{From: sender, To: receiver, Value: 100})
	})
	scheduler.Schedule(task)
	scheduler.Schedule(task)
	scheduler.Wait()

	time.Sleep(10 * time.Second)
	_sender, _ := bc.GetBalance(ctx, &pb.QueryParams{Address: sender})
	_receiver, _ := bc.GetBalance(ctx, &pb.QueryParams{Address: receiver})
	assert.Equal(t, float64(0), _sender.Balance)
	assert.Equal(t, float64(200), _receiver.Balance)
}
