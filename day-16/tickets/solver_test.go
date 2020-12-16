package tickets_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/advent-of-code/2020/day-16/tickets"
)

var _ = Describe("TicketScanner", func() {
	It("calculates the error rate of invalid fields", func() {
		subject := NewFieldScanner(fixtureNamed("1.txt"))

		Expect(subject.ErrorRate()).To(Equal(71))
	})
})

func fixtureNamed(file string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", file))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
