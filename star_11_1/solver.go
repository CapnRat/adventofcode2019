package star_11_1

import (
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_09_1"
	"strconv"
)

type Solver struct {}

const File = "star_11_1/input"

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Color int

const (
	Black Color = iota
	White
)

type GridPoint struct {
	x, y int
}

func (s *Solver) Solve() string {
	program := star_02_1.ReadProgramFromFile(File)

	grid := make(map[GridPoint]Color)
	var botPos GridPoint
	var botDir Direction

	input := make(chan int, 1)
	output := make(chan int)

	go func() {
		defer close(output)
		star_09_1.RunProgram(program, input, output)
	}()

	for {
		input <- int(grid[botPos])

		color, ok := <- output
		if !ok {
			break
		}

		grid[botPos] = Color(color)

		turn := <- output
		if turn == 0 {
			botDir--
		} else {
			botDir++
		}
		if botDir < Up {
			botDir = Left
		} else if botDir > Left {
			botDir = Up
		}

		switch botDir {
		case Up:
			botPos.y++
		case Right:
			botPos.x++
		case Down:
			botPos.y--
		case Left:
			botPos.x--
		}
	}

	return strconv.Itoa(len(grid))
}
