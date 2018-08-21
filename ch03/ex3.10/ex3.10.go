/*
	Exercise 3.10
*/

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	for i, v := range s {
		if i > 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
