/*
package chain is a library for creating blockchains.
	package main

	import (
		"fmt"

		"github.com/shilangyu/go-chain"
	)

	func main() {
		bc := chain.NewBlockChain()
		bc.AddBlock([]byte("Second block"))
		bc.AddBlock([]byte("Third block"))

		fmt.Printf("The blockchain is correct: %t", bc.Validate())

		bc[0].data = []byte("altered")

		fmt.Printf("The blockchain is correct: %t", bc.Validate())
	}
*/
package chain
