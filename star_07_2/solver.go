package star_07_2

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
	const phaseOffset = 5
	phases := [programChainLength]int{}
	bestPhase := [programChainLength]int{}
	maxThrust := 0
	for true {
		if !UpdatePhases(&phases, phaseHeight) {
			break
		}

		// one extra channel to listen to last -> first signals
		channels := make([]chan int, programChainLength + 1)
		for i, _ := range channels {
			channels[i] = make(chan int, 1)
		}
		buffer := 0
		for i := 0; i < programChainLength; i++ {
			programCopy := append([]int(nil), program...)
			go RunProgram(programCopy, channels[i], channels[i+1])
		}

		// send initial input get to get the whole thing started
		for i, phase := range phases {
			channels[i] <- phase + phaseOffset
		}
		channels[0] <- 0

		// capture the out of last program and loop it back
		for buffer = range channels[programChainLength] {
			channels[0] <- buffer
		}
		close (channels[0])
		if buffer > maxThrust {
			for i, phase := range phases {
				bestPhase[i] = phase + phaseOffset
			}
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

func RunProgram(program []int, input <-chan int, output chan<- int) {
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
			program[program[i+1]] = <- input
		case OpOut:
			value := program[i+1]
			if modes[0] == PositionMode {
				value = program[value]
			}
			output <- value
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
			close(output)
			return
		}

		i += length
	}

	panic("fail")
}