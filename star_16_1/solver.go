package star_16_1

import (
	"github.com/CapnRat/adventofcode2019/star_03_1"
	"io/ioutil"
	"strconv"
)

type Solver struct {}

const File = "star_16_1/input"

func (s *Solver) Solve() string {
	content, err := ioutil.ReadFile(File)
	if err != nil {
		panic("failed to read input")
	}

	return SolveInput(string(content))
}

func SolveInput(input string) string {
	return SolveInputWithPatternForPhases(input, []int{0,1,0,-1}, 100)
}

func SolveInputWithPatternForPhases(input string, pattern []int, phases int) string {
	signal := make([]int, len(input))
	for i, char := range input {
		signal[i], _ = strconv.Atoi(string(char))
	}

	for phase := 0; phase < phases; phase++ {
		newSignal := make([]int, len(signal))
		for i := 0; i < len(signal); i++ {
			var accum int
			for j := 0; j < len(signal); j++ {
				patternIdx := ((j + 1) / (i + 1)) % len(pattern)
				accum += signal[j] * pattern[patternIdx]
			}
			newSignal[i] = star_03_1.Abs(accum % 10)
		}
		signal = newSignal
	}

	var output string
	for i := 0; i < 8; i++ {
		output += strconv.Itoa(signal[i])
	}

	return output
}

