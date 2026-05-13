package popcount2

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(259)
	}
}

func BenchmarkPopCountByLookupViaLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLookupViaLoop(259)
	}
}

func BenchmarkPopCountBitByBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountBitByBit(259)
	}
}

func BenchmarkPopCountByClearingRightmostNonZeroBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearingRightmostNonZeroBit(259)
	}
}
