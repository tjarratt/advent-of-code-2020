package calculator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/tjarratt/advent-of-code/2020/day-18/calculator"
)

var _ = Describe("Bizarro Math Calculator", func() {
	It("disregards normal operator precedence", func() {
		Expect(calculator.Solve("1 + 2 * 3 + 4 * 5 + 6")).To(Equal(71))
	})

	It("handles parentheses with grace", func() {
		Expect(calculator.Solve("2 * 3 + (4 * 5)")).To(Equal(26))
		Expect(calculator.Solve("5 + (8 * 3 + 9 + 3 * 4 * 3)")).To(Equal(437))
	})

	It("handles nested parentheses with ease", func() {
		Expect(calculator.Solve("1 + (2 * 3) + (4 * (5 + 6))")).To(Equal(51))
		Expect(calculator.Solve("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")).To(Equal(12240))
		Expect(calculator.Solve("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")).To(Equal(13632))
	})
})
