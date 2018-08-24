/*
	Exercise 4.1
*/

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sum384 = flag.Bool("sha384", false, "use SHA384 to produce hash")
var sum512 = flag.Bool("sha512", false, "use SHA512 to produce hash")

func main() {
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		printHash([]byte(input.Text()))
	}
}

func printHash(data []byte) {
	if *sum384 {
		fmt.Printf("%x\n", sha512.Sum384(data))
	} else if *sum512 {
		fmt.Printf("%x\n", sha512.Sum512(data))
	} else {
		fmt.Printf("%x\n", sha256.Sum256(data))
	}
}
