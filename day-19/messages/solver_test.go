package messages_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/advent-of-code/2020/day-19/messages"
)

var _ = Describe("Detecting Corrupted Messages", func() {
	It("matches messages to rules", func() {
		solver := NewSolver(fixtureNamed("1.txt"))

		Expect(solver.UncorruptedMessages()).To(Equal(2))
	})

	It("detects messages that match the rules", func() {
		solver := NewSolver(fixtureNamed("2.txt"))

		Expect(solver.UncorruptedMessages()).To(Equal(2))
	})
})

func fixtureNamed(file string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", file))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
