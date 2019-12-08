package star_05_2

import (
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"os"
)

const file = "star_05_1/input"

const (
	OpNul = iota
	OpAdd
	OpMul
	OpIn
	OpOut
	OpJmpT
	OpJmpF
	OpLess
	OpEql
	OpHalt = 99
)
type OpCode int

const (
	PositionMode = iota
	ImmediateMode
)
type ParamMode int

type Solver struct{}

func (s *Solver) Solve () string {
	program := star_02_1.ReadProgramFromFile(file)

	RunProgram(program, 5)

	return "Program Done"
}

func ParseInstruction (instruction int) (opCode OpCode, modes []ParamMode, length int) {
	opCode = OpCode(instruction % 100)
	switch opCode {
	case OpNul:
		length = 1
	case OpAdd, OpMul:
		arg1Mode := ParamMode(instruction % 1000 / 100)
		arg2Mode := ParamMode(instruction % 10000 / 1000)
		modes = []ParamMode{arg1Mode, arg2Mode}
		length = 4
	case OpIn:
		length = 2
	case OpOut:
		arg1Mode := ParamMode(instruction % 1000 / 100)
		modes = []ParamMode{arg1Mode}
		length = 2
	case OpJmpT, OpJmpF:
		arg1Mode := ParamMode(instruction % 1000 / 100)
		arg2Mode := ParamMode(instruction % 10000 / 1000)
		modes = []ParamMode{arg1Mode, arg2Mode}
		length = 0
	case OpLess, OpEql:
		arg1Mode := ParamMode(instruction % 1000 / 100)
		arg2Mode := ParamMode(instruction % 10000 / 1000)
		modes = []ParamMode{arg1Mode, arg2Mode}
		length = 4
	case OpHalt:
		length = 1
	}
	return
}

func RunProgram(program []int, inputs ...int) {
	// input pointer
	inputPtr := 0
	// instruction pointer
	i := 0
	for true {
		instruction := program[i]
		op, modes, length := ParseInstruction(instruction)

		switch op {
		case OpAdd, OpMul, OpLess, OpEql:
			left := program[i+1]
			if modes[0] == PositionMode {
				left = program[left]
			}
			right := program[i+2]
			if modes[1] == PositionMode {
				right = program[right]
			}
			switch op {
			case OpAdd:
				program[program[i+3]] = left + right
			case OpMul:
				program[program[i+3]] = left * right
			case OpLess:
				value := 0
				if left < right { value = 1 }
				program[program[i+3]] = value
			case OpEql:
				value := 0
				if left == right { value = 1 }
				program[program[i+3]] = value
			}
		case OpIn:
			if len(inputs) > inputPtr {
				program[program[i+1]] = inputs[inputPtr]
				inputPtr++
			} else {
				fmt.Print("Input: ")
				in := 0
				_, err := fmt.Scanf("%d\n", &in)
				if err != nil {
					fmt.Println("failed to parse input")
				}
				program[program[i+1]] = in
			}
		case OpOut:
			out := program[i+1]
			if modes[0] == PositionMode {
				out = program[out]
			}
			fmt.Println(out)
		case OpJmpT, OpJmpF:
			test := program[i+1]
			if modes[0] == PositionMode {
				test = program[test]
			}
			value := program[i+2]
			if modes[1] == PositionMode {
				value = program[value]
			}
			if (op == OpJmpT && test != 0) || (op == OpJmpF && test == 0) {
				i = value
			} else {
				length = 3
			}
		case OpHalt:
			fmt.Println("Halting")
			return
		}

		i += length
	}

	fmt.Println("error unexpected end of program")
	os.Exit(1)
}