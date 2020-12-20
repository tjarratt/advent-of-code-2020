package tiles

import (
	"strconv"
	"strings"
)

func NewSolver(input string) solver {
	return solver{tiles: parse(input)}
}

type solver struct {
	tiles []tile
}

type tile struct {
	id   int
	data []string
}

func (s solver) Tiles() []tile {
	return s.tiles
}

func (s solver) Corners() []int {
	result := []int{}

	// strategy :: look for FOUR pieces each having two edges that do not match
	lookup := map[int][][]int{}
	cache := map[int]int{}

	for _, tile := range s.tiles {
		lookup[tile.id] = compute_borders(tile.data, cache)
	}

	for id, borders := range lookup {
		if is_corner(borders, cache) {
			result = append(result, id)
		}
	}

	return result
}

func is_corner(borders [][]int, cache map[int]int) bool {
	for _, border := range borders {
		sum := 0
		for _, v := range border {
			sum += cache[v]
		}

		if sum == 6 {
			return true
		}
	}

	return false
}

func compute_borders(data []string, cache map[int]int) [][]int {
	// top
	top := strings.Split(data[0], "")

	// bottom
	bottom := strings.Split(data[len(data)-1], "")

	left := []string{}
	right := []string{}
	for _, row := range data {
		left = append(left, string(row[0]))
		right = append(right, string(row[len(row)-1]))
	}

	cache[value_for(top)]++
	cache[value_for(bottom)]++
	cache[value_for(left)]++
	cache[value_for(right)]++
	cache[value_for(flip(top))]++
	cache[value_for(flip(bottom))]++
	cache[value_for(flip(left))]++
	cache[value_for(flip(right))]++

	return [][]int{
		[]int{value_for(top), value_for(bottom), value_for(left), value_for(right)},
		[]int{value_for(flip(top)), value_for(flip(bottom)), value_for(left), value_for(right)},
		[]int{value_for(top), value_for(bottom), value_for(flip(left)), value_for(flip(right))},
	}
}

func value_for(data []string) int {
	border := 0
	for index, char := range data {
		if char == "." {
			continue
		}

		border += pow(2, index)
	}

	return border
}

func flip(data []string) []string {
	result := make([]string, len(data))

	for i := 0; i < len(data); i++ {
		result[len(data)-1-i] = data[i]
	}

	return result
}

func pow(base, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= base
	}
	return result
}

// pragma mark - private
func parse(input string) []tile {
	result := []tile{}

	for _, chunk := range strings.Split(input, "\n\n") {
		lines := strings.Split(chunk, "\n")

		id, err := strconv.Atoi(strings.Split(strings.Split(lines[0], " ")[1], ":")[0])
		if err != nil {
			panic(err)
		}

		result = append(result, tile{id: id, data: nonempty(lines[1:])})
	}

	return result
}

func nonempty(lines []string) []string {
	result := []string{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		result = append(result, line)
	}

	return result
}
