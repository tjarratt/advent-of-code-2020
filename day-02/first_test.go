package main_test

import (
	"testing"

	. "github.com/tjarratt/advent-of-code-2020/day-02/rules"
)

func Test_is_valid(t *testing.T) {
	rule := Rule{Minimum: 1, Maximum: 3, Letter: "a"}
	password := "abcde"

	if !IsValid(password, rule) {
		t.Errorf("expected to be valid but is not (%s), %#v", password, rule)
	}
}

func Test_is_invalid(t *testing.T) {
	rule := Rule{Minimum: 1, Maximum: 3, Letter: "b"}
	password := "cdefg"

	if IsValid(password, rule) {
		t.Errorf("expected to be invalid but is not (%s), %#v", password, rule)
	}
}

func Test_is_valid_with_many_repeat(t *testing.T) {
	rule := Rule{Minimum: 2, Maximum: 9, Letter: "c"}
	password := "ccccccccc"

	if !IsValid(password, rule) {
		t.Errorf("expected to be valid but is not (%s), %#v", password, rule)
	}
}

func Test_is_not_valid_with_too_many_repeat(t *testing.T) {
	rule := Rule{Minimum: 2, Maximum: 9, Letter: "c"}
	password := "cccccccccc"

	if IsValid(password, rule) {
		t.Errorf("expected to be invalid but is not (%s), %#v", password, rule)
	}

}
