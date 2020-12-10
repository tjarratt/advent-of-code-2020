package jolted_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/advent-of-code/2020/day-10/jolted"
)

var _ = Describe("Joltage adapters", func() {
	It("determines the number of 1-jolt and 3-jolt differences in a chain of adapters", func() {
		solver := ChainedJoltageAdapters(fixtureNamed("1.txt"))

		Expect(solver.DifferencesOfJolts(1)).To(Equal(22))
		Expect(solver.DifferencesOfJolts(3)).To(Equal(10))
	})
})

func fixtureNamed(file string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", file))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
