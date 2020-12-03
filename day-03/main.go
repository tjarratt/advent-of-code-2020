package main

import (
	"io/ioutil"

	. "github.com/tjarratt/advent-of-code-2020/day-03/toboggan"
)

func main() {
	fixture, err := ioutil.ReadFile("input-1.txt")
	if err != nil {
		panic(err)
	}

	grid := NewGrid(fixture)
	println(grid.CountTrees(Trajectory{Right: 3, Down: 1}))
}
