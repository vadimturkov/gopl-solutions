/*
	Exercise 3.11
*/

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer

	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	var decPart string
	if dot := strings.Index(s, "."); dot >= 0 {
		decPart = s[dot:]
		s = s[:dot]
	}

	start := len(s) % 3
	if start == 0 {
		start = 3
	}

	buf.WriteString(s[:start])

	for i := start; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	if decPart != "" {
		buf.WriteString(decPart)
	}

	return buf.String()
}
