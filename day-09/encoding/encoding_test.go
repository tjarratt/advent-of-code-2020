package encoding_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/advent-of-code/2020/day-09/encoding"
)

var _ = Describe("Encoding Errors with XMAS", func() {
	It("determines the first number that is invalid", func() {
		solver := RingDecrypter(fixture("1.txt"), 5)

		Expect(solver.FirstInvalidNumber()).To(Equal(127))
	})
})

func fixture(name string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", name))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
