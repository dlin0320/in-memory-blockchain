# in-memory-blockchain

## Running
```bash
# running the server
make main
# running with docker
make docker
```

## Testing
```bash
# make sure the server is running then
make test
```

## Design details

### Server & Client
The server can be run locally or with docker, takes input from the client thru grpc API, and dumps the logs to the console.


### Blockchain
The blockchain is stored in a map for faster lookup time, and new blocks are periodically mined with a scheduler. Each block can contain up to 1000 transactions, and the transactions are checked twice to avoid race condition and double-spending:

1. when creating transactions
2. before mining a new block

The account balances are updated after mining a new block and stored as
```golang
type Balance struct {
    final float64
    temp float64
}
```
where __final__ is the actual balance and __temp__ is for handling multiple transactions from the same sender/receiver in a single block.
## References

https://go.dev/tour/welcome/1

https://gobyexample.com/

https://i0.wp.com/bitsonblocks.net/wp-content/uploads/2015/09/bitcoin_blockchain_infographic1.jpg?ssl=1

https://github.com/ethereum/go-ethereum/

https://grpc.io/docs/languages/go/quickstart/

https://github.com/bitcoin/bitcoin