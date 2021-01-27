package kata

import (
	"math/bits"
)

func CountBits(n uint) int {
	return bits.OnesCount(n)
}

func CountBits2(n uint) (cnt int) {
	for n > 0 {
		if n&1 == 1 {
			cnt += 1
		}
		n = n >> 1
	}
	return
}
