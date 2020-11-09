package star_16_2

import (
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

	return SolveInput(content)
}

func SolveInput(data []byte) string {

	offset, err := strconv.Atoi(string(data[:7]))
	if err != nil {
		panic(err)
	}

	length := (len(data) * 10000) - offset
	numRepeat := (length / len(data)) + 1
	var input []byte
	for i := 0; i < numRepeat; i++ {
		input = append(input, data...)
	}

	input = input[len(input) - length:]

	for i := 0; i < 100; i++ {
		var accum uint8 = 0
		for i := len(input) - 1; i >= 0; i-- {
			curValue, err := strconv.Atoi(string(input[i]))
			if err != nil {
				panic(err)
			}
			accum = (uint8(curValue) + accum) % 10
			input[i] = accum + '0'
		}
	}

	return string(input[:8])
}