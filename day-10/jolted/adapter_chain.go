package jolted

import (
	"sort"
	"strconv"
	"strings"
)

type adapterChain struct {
	adapters []adapter
}

type adapter struct {
	jolts int
}

func ChainedJoltageAdapters(raw string) adapterChain {
	return adapterChain{adapters: parse(raw)}
}

func (chain adapterChain) DifferencesOfJolts(quantity int) int {
	count := 0

	current := chain.adapters[0]
	for _, next := range chain.adapters[1:] {
		if (next.jolts - current.jolts) == quantity {
			count += 1
		}

		current = next
	}

	return count
}

// pragma mark - private
func parse(input string) []adapter {
	lines := strings.Split(input, "\n")
	charges := []int{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		charge, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		charges = append(charges, charge)
	}

	sort.Slice(charges, func(i, j int) bool {
		return charges[i] < charges[j]
	})

	// wall outlet has effective jolts of 0
	adapters := []adapter{adapter{jolts: 0}}

	for _, charge := range charges {
		adapter := adapter{jolts: charge}
		adapters = append(adapters, adapter)
	}

	// always add one for the handheld
	adapters = append(adapters, adapter{jolts: adapters[len(adapters)-1].jolts + 3})

	return adapters
}
