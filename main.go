package main

import (
	"github.com/dlin0320/in-memory-blockchain/server"
)

func main() {
	go server.Serve()
	server.Mine()
}
