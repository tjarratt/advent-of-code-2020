package tickets

import (
	"regexp"
	"strconv"
	"strings"
)

func NewFieldScanner(input string) solver {
	return parse(input)
}

type solver struct {
	rules []rule

	nearby_tickets [][]int
}

type rule struct {
	name     string
	validity fieldValidity
}

type fieldValidity struct {
	first  validRange
	second validRange
}

func (r rule) Valid(field int) bool {
	valid := false

	if field >= r.validity.first.min && field <= r.validity.first.max {
		valid = valid || true
	}
	if field >= r.validity.second.min && field <= r.validity.second.max {
		valid = valid || true
	}

	return valid
}

type validRange struct {
	min int
	max int
}

func (s solver) ErrorRate() int {
	result := 0

	for _, ticket := range s.nearby_tickets {
		for _, field := range ticket {
			any_valid := false
			for _, rule := range s.rules {
				if rule.Valid(field) {
					any_valid = true
					break
				}
			}

			if !any_valid {
				result += field
			}
		}
	}

	return result
}

func parse(input string) solver {
	rules := []rule{}
	nearby_tickets := [][]int{}

	sections := strings.Split(input, "\n\n")
	for _, line := range strings.Split(sections[0], "\n") {
		if len(line) == 0 {
			break
		}

		pieces := strings.Split(line, ":")
		rules = append(rules, rule{name: pieces[0], validity: parseValidity(pieces[1])})
	}

	for _, line := range strings.Split(sections[2], "\n")[1:] {
		if len(line) == 0 {
			break
		}

		ticket := []int{}
		for _, raw := range strings.Split(line, ",") {
			field, err := strconv.Atoi(raw)
			if err != nil {
				panic(err)
			}

			ticket = append(ticket, field)
		}
		nearby_tickets = append(nearby_tickets, ticket)
	}

	return solver{rules: rules, nearby_tickets: nearby_tickets}
}

func parseValidity(input string) fieldValidity {
	re := regexp.MustCompile("([0-9]+)-([0-9]+)")

	matches := re.FindAllString(input, -1)

	return fieldValidity{
		first:  mustParseRange(matches[0]),
		second: mustParseRange(matches[1]),
	}
}

func mustParseRange(raw string) validRange {
	pieces := strings.Split(raw, "-")

	min, err := strconv.Atoi(pieces[0])
	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(pieces[1])
	if err != nil {
		panic(err)
	}

	return validRange{min: min, max: max}
}
