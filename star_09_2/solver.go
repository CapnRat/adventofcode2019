package star_09_2

import (
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_09_1"
)

type Solver struct{}

func (s *Solver) Solve () string {
	program := star_02_1.ReadProgramFromFile(star_09_1.File)

	return star_09_1.SolveWithProgram(program, []int{2})
}