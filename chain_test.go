package chain

import (
	"testing"
)

func BenchmarkBlockChain(b *testing.B) {
	b.ReportAllocs()
	length := 10_000_000 - 1

	for n := 0; n < b.N; n++ {
		bc := NewBlockChain()

		for i := 0; i < length; i++ {
			bc.AddBlock([]byte(string(i)), 0)
		}
	}
}

func TestBlockChain_Validate(t *testing.T) {
	ok := NewBlockChain()
	ok.AddBlock([]byte("block"), 18)
	ok.AddBlock([]byte("block"), 18)

	empty := BlockChain{}

	altered := NewBlockChain()
	altered.AddBlock([]byte("block"), 0)
	altered.AddBlock([]byte("block"), 0)
	altered[1].Data = []byte("h#cked")

	missing := NewBlockChain()
	missing.AddBlock([]byte("block"), 0)
	missing.AddBlock([]byte("block"), 0)
	missing = append(missing[:1], missing[2:]...)

	noNonce := NewBlockChain()
	noNonce = append(noNonce, NewBlock(noNonce[0], []byte("block"), 10))

	tests := []struct {
		name string
		bc   BlockChain
		want bool
	}{
		{
			"Correct blockchain",
			ok,
			true,
		},
		{
			"Empty blockchain",
			empty,
			false,
		},
		{
			"Altered block",
			altered,
			false,
		},
		{
			"Missing block",
			missing,
			false,
		},
		{
			"Uncalculated nonce",
			noNonce,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.Validate(); got != tt.want {
				t.Errorf("BlockChain.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockChain_String(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock([]byte("block 2"), 0)

	tests := []struct {
		name string
		bc   BlockChain
		want string
	}{
		{
			"Bare",
			NewBlockChain(),
			"[GENESIS]",
		},
		{
			"2 blocks",
			bc,
			"[GENESIS -> block 2]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.String(); got != tt.want {
				t.Errorf("BlockChain.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
