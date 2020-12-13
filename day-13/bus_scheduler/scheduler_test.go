package bus_scheduler_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/advent-of-code/2020/day-13/bus_scheduler"
)

var _ = Describe("Convoluted Bus Schedules", func() {
	It("determines the first bus we could take, based on guesses", func() {
		solver := NewBusScheduleReader(fixtureNamed("1.txt"))

		Expect(solver.FirstBusAvailable()).To(Equal(59))
		Expect(solver.MinutesWaitingForBus()).To(Equal(5))
	})
})

func fixtureNamed(file string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", file))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
