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
	program := ReadProgramFromFile(file)

	// 1202 program alarm
	program[1] = 12
	program[2] = 2

	program = RunProgram(program)

	return strconv.Itoa(program[0])
}

func RunProgram(program []int) []int {
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

func ReadProgramFromFile(file string) []int {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	return ParseProgram(string(bytes))
}

func ParseProgram(input string) []int {
	instructions := strings.Split(input, ",")
	program := make([]int, len(input))
	for i, v := range instructions {
		instruction, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("error parsing program: ", v)
			os.Exit(1)
		}
		program[i] = instruction
	}
	return program
}