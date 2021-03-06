package chain

import (
	"bytes"
)

// BlockChain is a slice of Blocks
// it is recommended to initialize it from NewBlockChain, not manually
type BlockChain []*Block

// NewBlockChain initializes a BlockChain with a genesis block
func NewBlockChain() BlockChain {
	block := &Block{
		PrevHash:   []byte{},
		Data:       []byte("GENESIS"),
		Difficulty: 0,
	}
	block.FindNonce()
	block.GenerateHash()

	return append(BlockChain{}, block)
}

// AddBlock adds a new block to the chain
func (bc *BlockChain) AddBlock(data []byte, difficulty uint8) {
	block := NewBlock((*bc)[len(*bc)-1], data, difficulty)
	block.FindNonce()
	block.GenerateHash()

	*bc = append(*bc, block)
}

// String is a formated representation of a blockchain
func (bc BlockChain) String() string {
	s := "["

	for i, block := range bc {
		s += string(block.Data)
		if i != len(bc)-1 {
			s += " -> "
		}
	}

	s += "]"

	return s
}

// Validate makes sure the blockchain has no defects:
// Is the length is at least 1?
// Are the linked hashes between blocks correct?
// Is there a proof of work?
func (bc BlockChain) Validate() bool {
	correctHash := func(b Block) bool {
		presumedHash := b.Hash
		b.GenerateHash()
		return bytes.Equal(presumedHash, b.Hash)
	}

	if len(bc) < 1 || !correctHash(*bc[0]) {
		return false
	}

	for i, block := range bc[1:] {
		if !bytes.Equal(block.PrevHash, bc[i].Hash) || !correctHash(*block) || !block.ProofOfWork() {
			return false
		}
	}
	return true
}
