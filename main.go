package main

import (
	"github.com/dlin0320/in-memory-blockchain/server"
)

func main() {
	s := server.Serve()
	s.Mine()
}
