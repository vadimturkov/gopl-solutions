/*
	Exercise 4.1
*/

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%v", c1, c2, diff(c1, c2))
}

func diff(d1, d2 [32]byte) int {
	count := 0
	for i := range d1 {
		count += bitCount(d1[i], d2[i])
	}
	return count
}

func bitCount(byte1, byte2 byte) int {
	count := 0
	for i := uint8(0); i < 8; i++ {
		bit1 := (byte1 >> i) & 1
		bit2 := (byte2 >> i) & 1
		if bit1 != bit2 {
			count++
		}
	}
	return count
}
