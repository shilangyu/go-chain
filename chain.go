package main

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain is a slice of Blocks
// it is recommended to initialize it from NewBlockChain, not manually
type BlockChain []*Block

// NewBlockChain initializes a BlockChain with a genesis block
func NewBlockChain() BlockChain {
	block := &Block{
		prevHash: []byte{},
		data:     []byte("GENESIS"),
	}
	hash := sha256.Sum256(bytes.Join([][]byte{block.data, block.prevHash}, []byte{}))
	block.hash = hash[:]

	return append(BlockChain{}, block)
}

// AddBlock adds a new block to the chain
func (bc *BlockChain) AddBlock(data []byte) {
	*bc = append(*bc, NewBlock((*bc)[len(*bc)-1], data))
}

// String is a formated representation of a blockchain
func (bc BlockChain) String() string {
	s := "["

	for i, block := range bc {
		s += string(block.data)
		if i != len(bc)-1 {
			s += " -> "
		}
	}

	s += "]"

	return s
}
