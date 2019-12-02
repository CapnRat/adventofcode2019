package star_02_1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const file = "star_02_1/input"

type Solver struct{}

func (s *Solver) Solve () string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	input := strings.Split(string(bytes), ",")
	program := make([]int, len(input))
	for i, v := range input {
		program[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Println("error parsing input: ", v)
			os.Exit(1)
		}
	}

	// 1202 program alarm
	program[1] = 12
	program[2] = 2

	program = s.RunProgram(program)

	return strconv.Itoa(program[0])
}

func (s *Solver) RunProgram(program []int) []int {
	for i := 0; i < len(program); i += 4 {
		op := program[i]

		switch op {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		case 99:
			return program
		}
	}

	fmt.Println("error unexpected end of program")
	os.Exit(1)
	return nil
}