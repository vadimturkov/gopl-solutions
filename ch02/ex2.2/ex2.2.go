/*
	Exercise 2.2
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"./convert"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			process(arg)
		}
		return
	}

	fmt.Println("Please, input number and press Return")

	var arg string
	_, err := fmt.Scanf("%s", &arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "convert: %v\n", err)
		os.Exit(1)
	}
	process(arg)
}

func process(arg string) {
	val, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "convert: %v\n", err)
		os.Exit(1)
	}

	tempF := convert.Fahrenheit(val)
	tempC := convert.Celsius(val)
	fmt.Printf("%s = %s, %s = %s\n", tempF, convert.FToC(tempF), tempC, convert.CToF(tempC))

	lengthM := convert.Meter(val)
	lengthF := convert.Feet(val)
	fmt.Printf("%s = %s, %s = %s\n", lengthM, convert.MToF(lengthM), lengthF, convert.FToM(lengthF))

	weightK := convert.Kilogram(val)
	weightP := convert.Pound(val)
	fmt.Printf("%s = %s, %s = %s\n", weightK, convert.KToP(weightK), weightP, convert.PToK(weightP))
}
