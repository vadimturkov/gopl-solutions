/*
	Exercise 3.12
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Please provide two strings by arguments.")
		os.Exit(2)
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	if areAnagram(s1, s2) {
		fmt.Println("These strings are anagrams of each other.")
	} else {
		fmt.Println("These strings aren't anagrams of each other.")
	}

}

func areAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1 = sortStr(s1)
	s2 = sortStr(s2)

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func sortStr(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
