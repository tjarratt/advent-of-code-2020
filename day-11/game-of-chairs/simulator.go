package game_of_chairs

import "strings"

func NewSimulator(raw string) simulator {
	return simulator{parse(raw)}
}

type simulator struct {
	grid [][]place
}

type place int

const (
	void place = iota
	floor
	seat
	occupied
)

func (sim simulator) OccupiedOnceStable() int {
	displacements := -1
	for displacements != 0 {
		displacements = sim.performRound()

	}

	seats_occupied := 0
	for _, row := range sim.grid {
		for _, place := range row {
			if place == occupied {
				seats_occupied++
			}
		}
	}

	return seats_occupied
}

func (sim simulator) performRound() int {
	changes := []seatChange{}

	for y, row := range sim.grid {
		for x, place := range row {
			if place == void {
				continue
			}
			if place == floor {
				continue
			}

			neighbors := countNeighbors(sim.grid, x, y)

			if place == seat && neighbors == 0 {
				changes = append(changes, seatChange{x, y, occupied})
			} else if place == occupied && neighbors >= 4 {
				changes = append(changes, seatChange{x, y, seat})
			}
		}
	}

	for _, change := range changes {
		sim.grid[change.y][change.x] = change.state
	}

	return len(changes)
}

type seatChange struct {
	x     int
	y     int
	state place
}

// pragma mark - debug
func prettyprint(grid [][]place) {
	for _, row := range grid {
		for _, place := range row {
			if place == void {
				print(" ")
			}
			if place == floor {
				print(".")
			}
			if place == seat {
				print("L")
			}
			if place == occupied {
				print("#")
			}
		}
		println()
	}
}

// pragma mark - private

func parse(input string) [][]place {
	rows := [][]place{}

	for y, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			break
		}

		rows = append(rows, []place{void})

		for _, spot := range strings.Split(line, "") {
			if spot == "." {
				rows[y] = append(rows[y], floor)
			} else {
				rows[y] = append(rows[y], seat)
			}
		}

		rows[y] = append(rows[y], void)
	}

	width := len(rows[0])
	void_row := []place{}
	for i := 0; i < width; i++ {
		void_row = append(void_row, void)
	}

	rows = append([][]place{void_row}, rows...)
	rows = append(rows, void_row)

	return rows
}

func countNeighbors(grid [][]place, x, y int) int {
	neighbors := 0

	if grid[y][x-1] == occupied {
		neighbors++
	}
	if grid[y][x+1] == occupied {
		neighbors++
	}
	if grid[y-1][x] == occupied {
		neighbors++
	}
	if grid[y+1][x] == occupied {
		neighbors++
	}
	if grid[y-1][x+1] == occupied {
		neighbors++
	}
	if grid[y-1][x-1] == occupied {
		neighbors++
	}
	if grid[y+1][x+1] == occupied {
		neighbors++
	}
	if grid[y+1][x-1] == occupied {
		neighbors++
	}

	return neighbors
}
