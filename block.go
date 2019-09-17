package chain

import (
	"bytes"
	"crypto/sha256"
)

// Block is an element in a blockchain.
// It is recommended to initialize it from NewBlock, not manually.
type Block struct {
	// PrevHash is the hash to the previous block
	PrevHash []byte
	// Data stores the data block data
	Data []byte
	// Hash is the computed hash derived from PrevHash+Data
	Hash []byte
}

// NewBlock initializes a block derived from
// the previous block and given data.
func NewBlock(prevBlock *Block, data []byte) *Block {
	block := &Block{
		PrevHash: prevBlock.Hash,
		Data:     data,
	}
	block.GenerateHash()

	return block
}

// GenerateHash generates a hash for a given block.
func (b *Block) GenerateHash() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.PrevHash, b.Data}, []byte{}))
	b.Hash = hash[:]
}
