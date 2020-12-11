package main

import (
	"io/ioutil"

	. "github.com/tjarratt/advent-of-code/2020/day-11/game-of-chairs"
)

func main() {
	solver := NewSimulator(input())
	println(solver.OccupiedOnceStable())
}

func input() string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
