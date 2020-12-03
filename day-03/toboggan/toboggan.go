package toboggan

import (
	"fmt"
	"strings"
)

type Node int

const (
	Open = Node(iota)
	Tree
)

type Grid struct {
	Rows [][]Node
}

type Trajectory struct {
	Right int
	Down  int
}

func (t Trajectory) multiply(i int) Trajectory {
	return Trajectory{Right: t.Right * i, Down: t.Down * i}
}

func NewGrid(input []byte) Grid {
	rows := [][]Node{}
	lines := strings.Split(string(input), "\n")

	for index, line := range lines {
		rows = append(rows, []Node{})
		chars := strings.Split(line, "")
		for _, char := range chars {
			if strings.EqualFold(char, ".") {
				rows[index] = append(rows[index], Open)
			} else if strings.EqualFold(char, "#") {
				rows[index] = append(rows[index], Tree)
			} else {
				panic(fmt.Sprintf("attempt to create Grid from unknown input '%#v'", char))
			}
		}
	}

	return Grid{Rows: rows}
}

func (g Grid) Column(index int) []Node {
	result := []Node{}

	for _, line := range g.Rows {
		result = append(result, line[index])
	}

	return result
}

func (g Grid) Row(index int) []Node {
	return g.Rows[index]
}

func (g Grid) CountTrees(trajectory Trajectory) int {
	total := 0

	for i := 0; i < len(g.Rows); i += 1 {
		vec := trajectory.multiply(i)
		if vec.Down >= len(g.Rows) {
			break
		}

		node := g.Fetch(vec)
		if node == Tree {
			total += 1
		}
	}

	return total
}

func (g Grid) Fetch(trajectory Trajectory) Node {
	row := trajectory.Down
	column := trajectory.Right % len(g.Rows[0])

	return g.Rows[row][column]
}
