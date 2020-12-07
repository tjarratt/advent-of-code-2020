package luggage_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/advent-of-code/2020/day-07/luggage"
)

var _ = Describe("Bag Sorting", func() {
	It("determines which bags can include others", func() {
		sorter := NewBagSorter(fixtureNamed("1.txt"))

		shiny_bag_holders := sorter.BagsWhichCouldContain("shiny gold")

		Expect(shiny_bag_holders).To(HaveLen(4))
		Expect(shiny_bag_holders).To(ContainElements(
			"bright white", "muted yellow",
			"dark orange", "light red",
		))
	})
})

func fixtureNamed(name string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", name))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
