package star_17_2

import (
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_09_1"
	"github.com/CapnRat/adventofcode2019/star_17_1"
	"strconv"
)

type Solver struct {}

func (s *Solver) Solve() string {
	main  := "A,B,B,A,C,A,A,C,B,C\n"
	funcA := "R,8,L,12,R,8\n"
	funcB := "R,12,L,8,R,10\n"
	funcC := "R,8,L,8,L,8,R,8,R,10\n"

	inputStr := main + funcA + funcB + funcC + "n\n"

	program := star_02_1.ReadProgramFromFile(star_17_1.File)
	program[0] = 2

	input := make(chan int, 1)
	output := make(chan int)
	go func() {
		defer close(output)
		star_09_1.RunProgram(program, input, output)
	}()

	go func() {
		for _, r := range inputStr {
			input <- int(r)
		}
	}()

	last := 0
	for r := range output {
		fmt.Print(string(r))
		last = r
	}

	return strconv.Itoa(last)
}
