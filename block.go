package chain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"math"
	"math/big"
)

// Block is an element in a blockchain.
// It is recommended to initialize it from NewBlock, not manually.
type Block struct {
	// PrevHash is the hash to the previous block
	PrevHash []byte
	// Data stores the data block data
	Data []byte
	// Difficulty describes the # of leading 0s in ProofOfWork
	Difficulty uint8
	// Nonce is a guessed number to provide a PoW
	Nonce uint64
	// Hash is the computed hash derived from PrevHash+Data
	Hash []byte
}

// NewBlock initializes a block derived from
// the previous block and given data.
func NewBlock(prevBlock *Block, data []byte, difficulty uint8) *Block {
	block := &Block{
		PrevHash:   prevBlock.Hash,
		Data:       data,
		Difficulty: difficulty,
	}
	block.GenerateHash()

	return block
}

// GenerateHash generates a hash for a given block.
func (b *Block) GenerateHash() {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.BigEndian, b.Nonce)

	hash := sha256.Sum256(bytes.Join([][]byte{
		b.PrevHash,
		b.Data,
		[]byte{b.Difficulty},
		buff.Bytes(),
	}, []byte{}))
	b.Hash = hash[:]
}

// FindNonce finds a nounce for a given block with the set difficulty.
// Error is returned if no nonce was found.
func (b *Block) FindNonce() error {
	var hashHolder big.Int
	var nonce uint64
	var hash [32]byte
	threshold := big.NewInt(1)
	threshold.Lsh(threshold, 256-uint(b.Difficulty))

	salt := [][]byte{
		b.PrevHash,
		b.Data,
		[]byte{b.Difficulty},
	}

	for nonce < math.MaxUint64 {
		buff := new(bytes.Buffer)
		binary.Write(buff, binary.BigEndian, nonce)

		hash = sha256.Sum256(bytes.Join(salt, buff.Bytes()))
		hashHolder.SetBytes(hash[:])

		if hashHolder.Cmp(threshold) == -1 {
			b.Nonce = nonce
			b.Hash = hash[:]
			return nil
		}

		nonce++
	}

	return errors.New("no fitting nonce was found")
}
