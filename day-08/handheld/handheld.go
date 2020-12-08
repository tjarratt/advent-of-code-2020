package handheld

import (
	"fmt"
	"strconv"
	"strings"
)

type handheld struct {
	accumulator  int
	instructions []instruction
}

type instruction struct {
	operation opcode
	argument  int
}

type opcode int

const (
	acc opcode = iota
	jmp
	nop
)

func NewHandheld(input string) *handheld {
	return &handheld{
		accumulator:  0,
		instructions: parse(input),
	}
}

func (h *handheld) Run() {
	seen := map[int]bool{}

	for instruction_pointer := 0; seen[instruction_pointer] == false; {
		seen[instruction_pointer] = true

		instruction := h.instructions[instruction_pointer]

		switch instruction.operation {
		case jmp:
			instruction_pointer += instruction.argument
		case acc:
			h.accumulator += instruction.argument
			instruction_pointer += 1
		case nop:
			instruction_pointer += 1
		default:
			panic(fmt.Sprintf("unexpected instruction %#v", instruction.operation))
		}
	}
}

func (h *handheld) Accumulator() int {
	return h.accumulator
}

// pragma mark - private
func parse(input string) []instruction {
	instructions := []instruction{}

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		pieces := strings.Split(line, " ")
		op := pieces[0]
		argument, err := strconv.Atoi(pieces[1])
		if err != nil {
			panic(err)
		}

		var opcode opcode
		switch op {
		case "jmp":
			opcode = jmp
		case "acc":
			opcode = acc
		case "nop":
			opcode = nop
		}

		instructions = append(instructions, instruction{
			operation: opcode,
			argument:  argument,
		})
	}

	return instructions
}
