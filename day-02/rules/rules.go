package rules

import (
	"fmt"
	"strings"
)

type CheckedPassword struct {
	Password string
	Rule     Rule
}

type Rule struct {
	Letter  string
	Minimum int
	Maximum int
}

func IsValid(password string, rule Rule) bool {
	count := strings.Count(password, rule.Letter)

	if count < rule.Minimum {
		println("not enough", password, count, fmt.Sprintf("%#v", rule))
		return false
	}

	if count > rule.Maximum {
		println("too many", password, count, fmt.Sprintf("%#v", rule))
		return false
	}

	println("just enough", password, count, fmt.Sprintf("%#v", rule))
	return true
}
