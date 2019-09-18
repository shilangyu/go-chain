package chain

import (
	"fmt"
	"testing"
)

func TestBlock_String(t *testing.T) {
	block := NewBlock(&Block{}, []byte("block"), 12)
	block.FindNonce()
	block.GenerateHash()

	tests := []struct {
		name string
		b    Block
		want string
	}{
		{
			"standard block",
			*block,
			fmt.Sprintf("{prevHash: , data: block, nonce: %d, difficulty: 12, hash: %x}", block.Nonce, block.Hash),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Block.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
