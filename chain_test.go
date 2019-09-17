package main

import (
	"testing"
)

func BenchmarkBlockChain(b *testing.B) {
	b.ReportAllocs()
	length := 10_000_000 - 1

	for n := 0; n < b.N; n++ {
		bc := NewBlockChain()

		for i := 0; i < length; i++ {
			bc.AddBlock([]byte(string(i)))
		}
	}
}
