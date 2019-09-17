package main

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain is a slice of Blocks
// it is recommended to initialize it from NewBlockChain, not manually
type BlockChain []Block

// NewBlockChain initializes a BlockChain with a genesis block
func NewBlockChain() BlockChain {
	block := Block{
		prevHash: []byte{},
		data:     []byte("GENESIS"),
	}
	hash := sha256.Sum256(bytes.Join([][]byte{block.data, block.prevHash}, []byte{}))
	block.hash = hash[:]

	return append(BlockChain{}, block)
}
