package main

import (
	"io/ioutil"

	"github.com/tjarratt/advent-of-code-2020/day-19/messages"
)

func main() {
	solver := messages.NewSolver(input())

	println(solver.UncorruptedMessages())
}

func input() string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
