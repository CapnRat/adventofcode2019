package star_07_1

import (
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_05_2"
	"strconv"
)

const file = "star_07_1/input"

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

	return SolveWithProgram (program)
}

func SolveWithProgram(program []int) string {
	const programChainLength = 5
	const phaseHeight = 5
	phases := [programChainLength]int{}
	bestPhase := [programChainLength]int{}
	maxThrust := 0
	for true {
		if !UpdatePhases(&phases, phaseHeight) {
			break
		}
		buffer := 0
		for i := 0; i < programChainLength; i++ {
			programCopy := append([]int(nil), program...)
			buffer = RunProgram(programCopy, phases[i], buffer)
		}
		if buffer > maxThrust {
			bestPhase = phases
			maxThrust = buffer
		}
	}

	fmt.Println("bestPhase", bestPhase)

	return strconv.Itoa(maxThrust)
}

func UpdatePhases(phases *[5]int, height int) bool {
	phases[len(phases) - 1]++
	for i := len(phases) - 1; i >= 0; i-- {
		if phases[i] > height - 1 {
			if i == 0 {
				return false
			}
			phases[i] = 0
			phases[i - 1]++
		}
	}
	if !ValidatePhases (phases) {
		return UpdatePhases(phases, height)
	}
	return true
}

func ValidatePhases(phases *[5]int) bool {
	used := make(map[int]bool)
	for i := 0; i < len(phases); i++ {
		if _, ok := used[phases[i]]; ok {
			return false
		}
		used[phases[i]] = true
	}
	return true
}

func RunProgram(program []int, inputs ...int) (output int) {
	// input pointer
	inputPtr := 0
	// instruction pointer
	i := 0
	for true {
		instruction := program[i]
		op, modes, length := star_05_2.ParseInstruction(instruction)

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
			output = program[i+1]
			if modes[0] == PositionMode {
				output = program[output]
			}
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
			return
		}

		i += length
	}

	panic("fail")
}