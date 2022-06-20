package main

import (
	"github.com/dlin0320/in-memory-blockchain/server"
)

func main() {
	server.Serve()
	server.Mine()

}
