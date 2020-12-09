package encoding

import (
	"fmt"
	"strconv"
	"strings"
)

type ringDecrypter struct {
	preambleLength int
	input          []int
}

func RingDecrypter(raw string, preamble int) ringDecrypter {
	return ringDecrypter{input: parse(raw), preambleLength: preamble}
}

func (r ringDecrypter) FirstInvalidNumber() int {
	for i := r.preambleLength; i < len(r.input); i++ {
		found := false

		for j := i - r.preambleLength; j < i && !false; j++ {
			for k := j + 1; k < i; k++ {
				if r.input[j]+r.input[k] == r.input[i] {
					found = true
					break
				}
			}
		}

		if !found {
			return r.input[i]
		}
	}

	return -1
}

func parse(input string) []int {
	results := []int{}

	for i, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("invalid input on line %d :: %s", i, line))
		}

		results = append(results, number)
	}

	return results
}
