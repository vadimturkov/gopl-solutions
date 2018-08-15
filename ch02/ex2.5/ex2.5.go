/*
	Exercise 2.5:
	The expression x&(x-1) clears the rightmost non-zero bit of x.
	Write a version of PopCount that counts bits by using this fact,
	and assess its performance.
*/

package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func popCount2(x uint64) int {
	var result byte
	for i := byte(0); i < 8; i++ {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}

func popCount3(x uint64) int {
	var result byte
	for i := byte(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			result++
		}
	}
	return int(result)
}

func popCount4(x uint64) int {
	var result byte
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return int(result)
}
