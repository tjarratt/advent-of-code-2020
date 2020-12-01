package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var goal int = 2020

type intPair struct {
	first  int
	second int
}

func main() {
	stuff, err := ioutil.ReadFile("input-1.txt")
	if err != nil {
		panic(err)
	}

	contents := string(stuff)
	lines := strings.Split(contents, "\n")

	input := []int{}
	for i := 0; i < len(lines); i++ {
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			continue
		}

		input = append(input, num)
	}

	resultChan := make(chan (intPair), 1)

	for i := range input {
		go checkForDesiredResult(input[i], input, goal, resultChan)
	}

	result := <-resultChan

	println(result.first * result.second)
}

func checkForDesiredResult(given int, input []int, goal int, resultChan chan (intPair)) {
	for _, i := range input {
		if given+i != goal {
			continue
		}

		resultChan <- intPair{first: given, second: i}
		return
	}
}
