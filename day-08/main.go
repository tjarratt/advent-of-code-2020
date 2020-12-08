package main

import (
	"io/ioutil"

	"github.com/tjarratt/advent-of-code-2020/day-08/handheld"
)

func main() {
	handheld := handheld.NewHandheld(input())
	handheld.Run()

	println(handheld.Accumulator())
}

func input() string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
