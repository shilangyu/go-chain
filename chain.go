package main

import (
	"bytes"
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
	block.GenerateHash()

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

// Validate makes sure the blockchain has no defects
// checks if the length is at least 1
// checks for linked hashes between blocks
// checks for correctness of a hash in each block
func (bc BlockChain) Validate() bool {
	correctHash := func(b Block) bool {
		presumedHash := b.hash
		b.GenerateHash()
		return bytes.Equal(presumedHash, b.hash)
	}

	if len(bc) < 1 || !correctHash(*bc[0]) {
		return false
	}

	for i, block := range bc[1:] {
		if !bytes.Equal(block.prevHash, bc[i-1].hash) || !correctHash(*block) {
			return false
		}
	}
	return true
}
