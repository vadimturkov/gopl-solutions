/*
	Exercise 1.2
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("os.Args[%d] = %v\n", i+1, arg)
	}
}
