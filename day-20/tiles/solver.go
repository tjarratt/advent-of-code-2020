package tiles

import (
	"fmt"
	"strconv"
	"strings"
)

func NewSolver(input string) solver {
	return solver{tiles: parse(input)}
}

type solver struct {
	tiles []tile
	cache map[int]int
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
	s.cache = map[int]int{}

	for _, tile := range s.tiles {
		lookup[tile.id] = s.computeBorders(tile.data)
	}

	for id, borders := range lookup {
		if is_corner(borders, s.cache) {
			result = append(result, id)
		}
	}

	return result
}

func (s solver) Image(width, height int) string {

	// build up a list of tiles with edges and appearance count
	remainingPieces := s.parsePieces()

	solution := [][]puzzlePiece{{}}
	solvedPieces := map[int]puzzlePiece{}

	// pick a corner piece and place it first
	for _, piece := range remainingPieces {
		//		for index, orientation := range piece.orientations {
		if piece.orientations[0].isTopLeftCorner() {
			println("top left piece is", piece.id)
			solution[0] = append(solution[0], piece)
			break
		}
		//		}
	}

	topLeftCorner := solution[0][0]
	solvedPieces[topLeftCorner.id] = topLeftCorner
	delete(remainingPieces, topLeftCorner.id)

	// STRATEGY :: build the border
	// look for a piece where an edge equals one of the edges we have
	// place it such that its edge with frequency 1 is OUTSIDE
	//
	// TACTICALLY build the top row

	prev := topLeftCorner
	for i := 0; i < 3; i++ {
		println("looking for a piece to connect to", prev.id, "on edge", prev.orientations[0].edges[3].String())
		for _, piece := range remainingPieces {
			var next puzzlePiece

			for index, orientation := range piece.orientations {
				if orientation.edges[2] == prev.orientations[0].edges[3] {
					next = piece.choose(index)
					break
				}
			}

			if next.id != 0 {
				solution[0] = append(solution[0], next)
				solvedPieces[next.id] = next
				delete(remainingPieces, next.id)
				prev = next
				break
			}
		}
	}

	println(fmt.Sprintf("the solution thus far ... (%d pieces solved) (%d remaining)", len(solvedPieces), len(remainingPieces)))
	println()

	// for each row of tiles
	for y_tile := 0; y_tile < len(solution); y_tile++ {
		// for each of the rows over in that image
		for y_row := 0; y_row < len(solution[y_tile][0].data); y_row++ {
			// for each of the tiles in this slice
			for _, tile := range solution[y_tile] {
				line := tile.data[y_row]
				print(line)
			}
			println()
		}
	}
	println()

	return "totoro"
}

type puzzlePiece struct {
	id           int
	data         []string
	orientations []orientation
}

func (p puzzlePiece) choose(index int) puzzlePiece {
	if index == 0 {
		println("chose tile", p.id, "with identity transformation")
		return p
	} else if index == 1 {
		// rotate left
		println("chose tile", p.id, "with left rotation")
		return puzzlePiece{
			id:           p.id,
			data:         rotateLeft(p.data),
			orientations: []orientation{p.orientations[1]},
		}
	} else if index == 2 {
		// rotate right
		println("chose tile", p.id, "with right rotation")
		return puzzlePiece{
			id:           p.id,
			data:         rotateRight(p.data),
			orientations: []orientation{p.orientations[2]},
		}
	} else if index == 3 {
		// rotate 180
		println("chose tile", p.id, "with rotation by 180 degrees")
		return puzzlePiece{
			id:           p.id,
			data:         rotate180(p.data),
			orientations: []orientation{p.orientations[3]},
		}
	} else {
		panic(fmt.Sprintf("Unexpected orientation: %d", index))
	}
}

func rotate180(data []string) []string {
	grid := make([][]string, len(data))
	for y, line := range data {
		for _, str := range strings.Split(line, "") {
			grid[y] = append(grid[y], str)
		}
	}

	// everything gets flipped
	rotated := make([]string, len(data))
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			x_index := len(grid[0]) - 1 - x
			y_index := len(grid) - 1 - y
			rotated[y] += grid[y_index][x_index]
		}
	}

	return rotated
}

func rotateLeft(data []string) []string {
	grid := make([][]string, len(data))
	for y, line := range data {
		for _, str := range strings.Split(line, "") {
			grid[y] = append(grid[y], str)
		}
	}

	rotated := make([]string, len(data))
	for x := len(grid[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(grid); y++ {
			rotated[y] += grid[y][x]
		}
	}

	println("rotating data left ??")
	println("original:")
	for _, line := range data {
		println(line)
	}

	println()
	println("rotated:")
	for _, line := range rotated {
		println(line)
	}
	println()

	return rotated
}

func rotateRight(data []string) []string {
	grid := make([][]string, len(data))
	for y, line := range data {
		for _, str := range strings.Split(line, "") {
			grid[y] = append(grid[y], str)
		}
	}

	rotated := make([]string, len(data))
	for x := 0; x < len(grid[0]); x++ {
		for y := len(grid) - 1; y >= 0; y++ {
			rotated[len(grid)-1-y] = grid[y][x]
		}
	}

	println("rotating data  ??")
	println("original:")
	for _, line := range data {
		println(line)
	}

	println()
	println("rotated:")
	for _, line := range rotated {
		println(line)
	}
	println()

	return rotated
}

type orientation struct {
	edges []edge
}

type edge struct {
	data      string
	frequency int
}

func (e edge) String() string {
	return strconv.Itoa(value_for(strings.Split(e.data, "")))
}

func (o orientation) isTopLeftCorner() bool {
	return o.isTopBorder() && o.isLeftBorder()
}

func (o orientation) isTopBorder() bool {
	return o.edges[0].frequency == 1
}

func (o orientation) isLeftBorder() bool {
	return o.edges[2].frequency == 1
}

func (s solver) parsePieces() map[int]puzzlePiece {
	result := map[int]puzzlePiece{}

	lookup := map[int][][]int{}
	s.cache = map[int]int{}

	for _, tile := range s.tiles {
		lookup[tile.id] = s.computeBorders(tile.data)
	}

	for _, tile := range s.tiles {
		result[tile.id] = puzzlePiece{
			id:           tile.id,
			data:         tile.data,
			orientations: s.orientationsFor(tile.data),
		}
	}

	return result
}

func (s solver) orientationsFor(data []string) []orientation {
	return []orientation{
		// 0 identity
		{
			edges: []edge{
				s.topSlice(data), s.bottomSlice(data),
				s.leftSlice(data), s.rightSlice(data),
			},
		},
		// 1. rotate left (anti-clockwise)
		{
			edges: []edge{
				s.rightSlice(data), s.leftSlice(data),
				s.flip(s.topSlice(data)), s.flip(s.bottomSlice(data)),
			},
		},
		// 2. rotate right (clockwise)
		{
			edges: []edge{
				s.flip(s.leftSlice(data)), s.flip(s.rightSlice(data)),
				s.bottomSlice(data), s.topSlice(data),
			},
		},
		// 3. rotate 180 degrees
		{
			edges: []edge{
				s.flip(s.bottomSlice(data)), s.flip(s.topSlice(data)),
				s.flip(s.rightSlice(data)), s.flip(s.leftSlice(data)),
			},
		},
	}
}

func (s solver) flip(e edge) edge {
	data := make([]string, len(e.data))
	for i, char := range e.data {
		data[len(data)-i-1] = string(char)
	}

	return edge{data: strings.Join(data, ""), frequency: s.cache[value_for(data)]}
}

// top and bottom always read their values from left to right
func (s solver) topSlice(data []string) edge {
	return edge{
		data:      data[0],
		frequency: s.cache[value_for(strings.Split(data[0], ""))],
	}
}

func (s solver) bottomSlice(data []string) edge {
	return edge{
		data:      data[len(data)-1],
		frequency: s.cache[value_for(strings.Split(data[len(data)-1], ""))],
	}
}

// left and right always read their values from top to bottomm
func (s solver) leftSlice(data []string) edge {
	slice := []string{}
	for index := 0; index < len(data); index++ {
		slice = append(slice, strings.Split(data[index], "")[0])
	}

	return edge{
		data:      strings.Join(slice, ""),
		frequency: s.cache[value_for(slice)],
	}
}

func (s solver) rightSlice(data []string) edge {
	slice := []string{}
	for index := 0; index < len(data); index++ {
		pieces := strings.Split(data[index], "")
		slice = append(slice, pieces[len(pieces)-1])
	}

	return edge{
		data:      strings.Join(slice, ""),
		frequency: s.cache[value_for(slice)],
	}
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

func (s solver) computeBorders(data []string) [][]int {
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

	s.cache[value_for(top)]++
	s.cache[value_for(bottom)]++
	s.cache[value_for(left)]++
	s.cache[value_for(right)]++
	s.cache[value_for(flip(top))]++
	s.cache[value_for(flip(bottom))]++
	s.cache[value_for(flip(left))]++
	s.cache[value_for(flip(right))]++

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