package star_13_1

import (
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_09_1"
	"strconv"
)

type Solver struct {}

const File = "star_13_1/input"

type TileID int

const (
	Empty TileID = iota
	Wall
	Block
	HorizPaddle
	Ball
)

type Position struct {
	x, y int
}

func (s *Solver) Solve() string {
	program := star_02_1.ReadProgramFromFile(File)

	input := make(chan int, 1)
	output := make(chan int)
	go func() {
		defer close(output)
		star_09_1.RunProgram(program, input, output)
	}()

	var maxX, maxY int
	tiles := make(map[Position]TileID)
	for {
		x, ok := <- output
		if !ok {
			break
		}
		y  := <- output
		id := <- output

		tiles[Position{x,y}] = TileID(id)

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	var numBlocks int
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			switch tiles[Position{x, y}] {
			case Empty:
				fmt.Print(" ")
			case Wall:
				fmt.Print("#")
			case Block:
				fmt.Print("▬")
				numBlocks++
			case HorizPaddle:
				fmt.Print("-")
			case Ball:
				fmt.Print("O")
			}
		}
		fmt.Println()
	}

	return strconv.Itoa(numBlocks)
}

