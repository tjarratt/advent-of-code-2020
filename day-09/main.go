package main

import (
	"io/ioutil"

	"github.com/tjarratt/advent-of-code-2020/day-09/encoding"
)

func main() {
	solver := encoding.RingDecrypter(input(), 25)

	println(solver.FirstInvalidNumber())
}

func input() string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
