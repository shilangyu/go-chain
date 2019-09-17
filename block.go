package main

import (
	"bytes"
	"crypto/sha256"
)

// Block is an element in a blockchain
// it is recommended to initialize it from NewBlock, not manually
type Block struct {
	prevHash []byte
	data     []byte
	hash     []byte
}

// NewBlock initializes a block derived from
// the previous block and given data
func NewBlock(prevBlock *Block, data []byte) *Block {
	block := &Block{
		prevHash: prevBlock.hash,
		data:     data,
	}
	block.GenerateHash()

	return block
}

// GenerateHash generates a hash for a given block
func (b *Block) GenerateHash() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.prevHash, b.data}, []byte{}))
	b.hash = hash[:]
}
