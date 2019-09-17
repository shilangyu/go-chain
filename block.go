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
func NewBlock(prevBlock Block, data []byte) Block {
	hash := sha256.Sum256(bytes.Join([][]byte{prevBlock.prevHash, data}, []byte{}))

	return Block{
		prevBlock.hash,
		data,
		hash[:],
	}
}
