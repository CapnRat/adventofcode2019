package star_14_2

import (
	"github.com/CapnRat/adventofcode2019/star_14_1"
	"strconv"
)

type Solver struct {}

func (s *Solver) Solve() string {
	return strconv.Itoa(SolveFromInput(star_14_1.File))
}

func SolveFromInput(path string) int {
	recipes := star_14_1.ReadRecipesFromInput(path)

	increment := 1000000000000
	for fuel := 0;; fuel += increment {
		ore := star_14_1.CalculateOreForFuel(recipes, fuel)
		if ore > 1000000000000 {
			fuel -= increment
			if increment == 1 {
				return fuel
			} else {
				increment /= 10
			}
		}
	}
}

