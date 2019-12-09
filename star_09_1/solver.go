package star_09_1

import (
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"strconv"
	"strings"
)

const File = "star_09_1/input"

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
	OpRBOff
	OpHalt = 99
)
type OpCode int

const (
	PositionMode = iota
	ImmediateMode
	RelativeMode
)
type ParamMode int

type Solver struct{}

func (s *Solver) Solve () string {
	program := star_02_1.ReadProgramFromFile(File)

	return SolveWithProgram(program, []int{1})
}

func SolveWithProgram(program []int, args []int) string {
	input := make(chan int)
	output := make(chan int)

	go func() {
		defer close(output)
		RunProgram(program, input, output)
	}()

	// initial input to get things rolling
	go func() {
		defer close(input)
		for _, arg := range args {
			input <- arg
		}
	}()

	var result []string
	// capture the out of last program and loop it back
	for buffer := range output {
		result = append(result, strconv.Itoa(buffer))
	}

	return strings.Join(result, ",")
}

func ParseInstruction (instruction int) (opCode OpCode, modes []ParamMode, length int) {
	opCode = OpCode(instruction % 100)
	switch opCode {
	case OpNul, OpHalt:
		length = 1
	case OpAdd, OpMul, OpLess, OpEql:
		arg1Mode := ParamMode(instruction % 1000 / 100)
		arg2Mode := ParamMode(instruction % 10000 / 1000)
		arg3Mode := ParamMode(instruction % 100000 / 10000)
		modes = []ParamMode{arg1Mode, arg2Mode, arg3Mode}
		length = 4
	case OpJmpT, OpJmpF:
		arg1Mode := ParamMode(instruction % 1000 / 100)
		arg2Mode := ParamMode(instruction % 10000 / 1000)
		modes = []ParamMode{arg1Mode, arg2Mode}
		length = 3
	case OpIn, OpOut, OpRBOff:
		arg1Mode := ParamMode(instruction % 1000 / 100)
		modes = []ParamMode{arg1Mode}
		length = 2
	}
	return
}

func RunProgram(program []int, input <-chan int, output chan<- int) {
	// load program into memory
	memory := make(map[int]int)
	for i, instruction := range program {
		memory[i] = instruction
	}

	rb := 0 // relative base
	i := 0 // instruction pointer
	for true {
		instruction := memory[i]
		op, modes, length := ParseInstruction(instruction)

		switch op {
		case OpAdd, OpMul, OpLess, OpEql:
			left := memory[i+1]
			if modes[0] == PositionMode {
				left = memory[left]
			} else if modes[0] == RelativeMode {
				left = memory[rb + left]
			}
			right := memory[i+2]
			if modes[1] == PositionMode {
				right = memory[right]
			} else if modes[1] == RelativeMode {
				right = memory[rb + right]
			}
			ptr := memory[i+3]
			if modes[2] == RelativeMode {
				ptr += rb
			}
			switch op {
			case OpAdd:
				memory[ptr] = left + right
			case OpMul:
				memory[ptr] = left * right
			case OpLess:
				value := 0
				if left < right { value = 1 }
				memory[ptr] = value
			case OpEql:
				value := 0
				if left == right { value = 1 }
				memory[ptr] = value
			}
		case OpIn:
			ptr := memory[i+1]
			if modes[0] == RelativeMode {
				ptr += rb
			}
			memory[ptr] = <- input
		case OpOut:
			value := memory[i+1]
			if modes[0] == PositionMode {
				value = memory[value]
			} else if modes[0] == RelativeMode {
				value = memory[rb + value]
			}
			output <- value
		case OpJmpT, OpJmpF:
			test := memory[i+1]
			if modes[0] == PositionMode {
				test = memory[test]
			} else if modes[0] == RelativeMode {
				test = memory[rb + test]
			}
			value := memory[i+2]
			if modes[1] == PositionMode {
				value = memory[value]
			} else if modes[1] == RelativeMode {
				value = memory[rb + value]
			}
			if (op == OpJmpT && test != 0) || (op == OpJmpF && test == 0) {
				i = value
				length = 0
			}
		case OpRBOff:
			offset := memory[i+1]
			if modes[0] == PositionMode {
				offset = memory[offset]
			} else if modes[0] == RelativeMode {
				offset = memory[rb + offset]
			}
			rb += offset
		case OpHalt:
			return
		}

		i += length
	}

	panic("fail")
}