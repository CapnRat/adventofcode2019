package star_11_2

import (
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_09_1"
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

	grid[botPos] = White

	var minx, miny, maxx, maxy int

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
			if botPos.y > maxy {
				maxy = botPos.y
			}
		case Right:
			botPos.x++
			if botPos.x > maxx {
				maxx = botPos.x
			}
		case Down:
			botPos.y--
			if botPos.y < miny {
				miny = botPos.y
			}
		case Left:
			botPos.x--
			if botPos.x < minx {
				minx = botPos.x
			}
		}
	}

	for y := maxy; y >= miny; y-- {
		for x := minx; x <= maxx; x++ {
			color := grid[GridPoint{x,y}]
			switch color {
			case Black:
				fmt.Print(" ")
			case White:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}

	return "Done"
}
