package main

import (
	"fmt"
)

func main() {
	blockchain := NewBlockChain()
	blockchain = append(blockchain, NewBlock(blockchain[0], []byte("second block")))

	fmt.Println(blockchain)
}
