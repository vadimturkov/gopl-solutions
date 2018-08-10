/*
	Exercise 1.3:
	Experiment to measure the difference in running time between our potentially
	inefficient versions and the one that uses strings.Join.
*/

package main

import (
	"os"
	"strings"
)

func echo1(args []string) string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2(args []string) string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3(args []string) string {
	return strings.Join(os.Args[1:], " ")
}
