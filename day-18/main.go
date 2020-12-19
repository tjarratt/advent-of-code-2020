package main

import (
	"io/ioutil"
	"strings"

	"github.com/tjarratt/advent-of-code-2020/day-18/calculator"
)

func main() {
	result := 0

	for _, line := range strings.Split(input(), "\n") {
		if len(line) == 0 {
			continue
		}

		result += calculator.Solve(line)
	}

	println(result)
}

func input() string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
