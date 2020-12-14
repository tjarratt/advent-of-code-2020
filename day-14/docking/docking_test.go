package docking_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/advent-of-code/2020/day-14/docking"
)

var _ = Describe("Docking protocol", func() {
	It("sums the values in memory, staying mindful of the mask", func() {
		solver := NewDockingProgram(fixtureNamed("1.txt"))

		Expect(solver.Solution()).To(Equal(165))
	})
})

func fixtureNamed(file string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", file))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
