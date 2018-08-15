package popcount

import (
	"testing"
)

const number uint64 = 0x1234567890ABCDEF

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount1(number)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount2(number)
	}
}
