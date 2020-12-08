package handheld_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/advent-of-code/2020/day-08/handheld"
)

var _ = Describe("Handheld Computers", func() {
	It("reports its accumulator when it reaches an infinite loop", func() {
		handheld := NewHandheld(fixtureNamed("1.txt"))

		handheld.Run()

		Expect(handheld.Accumulator()).To(Equal(5))
	})
})

func fixtureNamed(file string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", file))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
