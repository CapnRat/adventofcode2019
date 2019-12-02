package star_02_2

import (
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"strconv"
)

const file = "star_02_1/input"

type Solver struct{}

func (s *Solver) Solve () string {
	program := star_02_1.ReadProgramFromFile(file)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			candidate := append([]int(nil), program...)
			candidate[1] = noun
			candidate[2] = verb

			result := star_02_1.RunProgram(candidate)[0]

			// doubt this number is the same for everyone's puzzle
			if result == 19690720 {
				return strconv.Itoa(100 * noun + verb)
			}
		}
	}

	fmt.Println("error unexpected end of program")
	return ""
}