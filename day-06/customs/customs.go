package customs

import "strings"

type customsCounter struct {
	groups []Group
}

type Group struct {
	answers map[string]bool
}

func CustomsCounter(input string) customsCounter {
	return customsCounter{groups: parse(input)}
}

func (counter customsCounter) Sum() int {
	sum := 0

	for _, group := range counter.groups {
		sum += len(group.answers)
	}

	return sum
}

func (counter customsCounter) Groups() []Group {
	return counter.groups
}

func (group Group) Responses() []string {
	keys := make([]string, 0, len(group.answers))

	for key := range group.answers {
		keys = append(keys, key)
	}

	return keys
}

// pragma mark - private
func parse(input string) []Group {
	groups := []Group{}
	answers := map[string]bool{}

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			groups = append(groups, Group{answers: answers})
			answers = map[string]bool{}
		}

		for _, answer := range strings.Split(line, "") {
			answers[answer] = true
		}
	}

	return groups
}
