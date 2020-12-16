package main

import (
	"io/ioutil"

	"github.com/tjarratt/advent-of-code-2020/day-16/tickets"
)

func main() {
	solver := tickets.NewFieldScanner(input())

	println(solver.ErrorRate())
}

func input() string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
