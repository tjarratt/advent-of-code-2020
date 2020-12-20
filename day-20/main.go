package main

import (
	"io/ioutil"

	"github.com/tjarratt/advent-of-code-2020/day-20/tiles"
)

func main() {
	solver := tiles.NewSolver(input())

	product := 1
	for _, id := range solver.Corners() {
		product *= id
	}

	println(product)
}

func input() string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
