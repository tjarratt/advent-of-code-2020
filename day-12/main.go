package main

import (
	"io/ioutil"

	. "github.com/tjarratt/advent-of-code-2020/day-12/navigation"
)

func main() {
	navigator := NavigationAssistant(input())

	println(navigator.ManhattanDistance())
}

func input() string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
