package star_13_2

import (
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_09_1"
	"github.com/buger/goterm"
	"strconv"
	"time"
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

func (s *Solver) Solve() string {
	program := star_02_1.ReadProgramFromFile(File)

	// play for free!
	program[0] = 2

	input := make(chan int, 1)
	output := make(chan int)
	go func() {
		defer close(output)
		star_09_1.RunProgram(program, input, output)
	}()

	goterm.Clear()
	var score, paddleX, ballX, maxY int
	for {
		x, ok := <- output
		if !ok {
			return strconv.Itoa(score)
		}
		y  := <- output
		id := <- output

		if x == -1 && y == 0 {
			score = id
			goterm.MoveCursor(1, 1)
			goterm.Println(score)
			goterm.MoveCursor(1, maxY + 3)
			goterm.Flush()
			continue
		}

		goterm.MoveCursor(x + 1, y + 2)

		switch TileID(id) {
		case Empty:
			goterm.Print(" ")
		case Wall:
			goterm.Print("#")
		case Block:
			goterm.Print("▬")
		case HorizPaddle:
			goterm.Print("-")
			paddleX = x
		case Ball:
			goterm.Print("O")
			ballDir := x - ballX
			ballX = x
			if ballX + ballDir < paddleX {
				input <- -1
			} else if ballX + ballDir > paddleX + 1 {
				input <- 1
			} else {
				input <- 0
			}
		}
		if y > maxY {
			maxY = y
		}
		goterm.MoveCursor(1, maxY + 3)
		goterm.Flush()

		if TileID(id) == HorizPaddle {
			time.Sleep(10 * time.Millisecond)
		}
	}
}

