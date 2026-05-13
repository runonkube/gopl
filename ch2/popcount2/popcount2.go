package popcount2

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByLookupViaLoop(x uint64) int {
	sum := byte(0)
	for index := 0; index < 8; index++ {
		sum += pc[byte(x>>(index*8))]
	}
	return int(sum)
}

func PopCountBitByBit(x uint64) int {
	bitsSet := 0
	curr := x
	for index := 0; index < 64; index++ {
		if curr&1 == 1 {
			bitsSet++
		}
		curr = (curr >> 1)
	}
	return bitsSet
}

func PopCountByClearingRightmostNonZeroBit(x uint64) int {
	bitsSet := 0
	curr := x
	for curr > 0 {
		bitsSet++
		curr = curr & (curr - 1)
	}
	return bitsSet
}
