package boarding_passes

import (
	"fmt"
	"math"
	"strings"
)

type Seat struct {
	Row    int
	Column int
}

type boardingPass struct {
	raw string
}

func BoardingPass(input string) boardingPass {
	return boardingPass{raw: input}
}

func (pass boardingPass) Location() Seat {
	return Seat{
		Row:    parseRow(pass.raw[0:7]),
		Column: parseColumn(pass.raw[7:10]),
	}
}

func (pass boardingPass) SeatId() int {
	loc := pass.Location()

	return loc.Row*8 + loc.Column
}

// pragma mark - private
func parseRow(commands string) int {
	min := 0
	max := 127
	mid := 0

	for _, cmd := range strings.Split(commands, "") {
		switch cmd {
		case "B": // upper
			mid = int(math.Ceil(float64(min+max) / 2.0))
			min = mid
		case "F": // lower
			mid = int(math.Floor(float64(min+max) / 2.0))
			max = mid
		default:
			panic(fmt.Sprintf("unknown row command '%s'", cmd))
		}
	}

	return mid
}

func parseColumn(commands string) int {
	min := 0
	max := 7
	mid := 0

	for _, cmd := range strings.Split(commands, "") {
		switch cmd {
		case "L": // lower half
			mid = int(math.Floor(float64(min+max) / 2.0))
			max = mid

		case "R": // upper half
			mid = int(math.Ceil(float64(min+max) / 2.0))
			min = mid

		default:
			panic(fmt.Sprintf("unknown column command '%s'", cmd))
		}
	}

	return int(mid)
}
