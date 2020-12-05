package main

import (
	"io/ioutil"
	"strings"

	. "github.com/tjarratt/advent-of-code-2020/day-05/boarding-passes"
)

func main() {
	max_seat_id := -1

	for _, line := range readInput() {
		pass := BoardingPass(line)
		seatId := pass.SeatId()

		if seatId > max_seat_id {
			max_seat_id = seatId
		}
	}

	println(max_seat_id)
}

func readInput() []string {
	stuff, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(stuff), "\n")

	return lines[:len(lines)-1]
}
