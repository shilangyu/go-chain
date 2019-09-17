package main

import (
	"fmt"
)

func main() {
	blockchain := NewBlockChain()
	blockchain.AddBlock([]byte("second block"))

	fmt.Println(blockchain)
}
