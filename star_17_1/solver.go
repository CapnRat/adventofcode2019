package star_17_1

import (
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_09_1"
	"github.com/CapnRat/adventofcode2019/star_15_1"
	"github.com/buger/goterm"
)

type Solver struct {}

const File = "star_17_1/input"

type GridID int

const (
	Scaffold GridID = iota
	Empty
	Intersection
)

type Position star_15_1.Position

type Grid map[Position]GridID

func (s *Solver) Solve() string {
	_, output := ReadAndRunProgram(File)

	grid := make(Grid)

	var botPos Position
	var botDir star_15_1.Direction
	var curPos Position
	var maxPos Position
	goterm.Clear()
	for ascii := range output {
		if curPos.X > maxPos.X {
			maxPos.X = curPos.X
		}
		if curPos.Y > maxPos.Y {
			maxPos.Y = curPos.Y
		}

		switch ascii {
		case '#':
			grid[curPos] = Scaffold
		case '.':
			grid[curPos] = Empty
		case '^':
			grid[curPos] = Scaffold
			botPos = curPos
			botDir = star_15_1.North
		case '<':
			grid[curPos] = Scaffold
			botPos = curPos
			botDir = star_15_1.West
		case '>':
			grid[curPos] = Scaffold
			botPos = curPos
			botDir = star_15_1.East
		case 'v':
			grid[curPos] = Scaffold
			botPos = curPos
			botDir = star_15_1.South
		case 10:
			curPos.Y++
			curPos.X = -1
		default:
			panic(fmt.Errorf("unexpected intcode %d", ascii))
		}
		curPos.X++
	}

	calib := 0
	for pos, id := range grid {
		if id != Scaffold {
			continue
		}

		if pos.X == 0 || pos.Y == 0 || pos.X == maxPos.X || pos.Y == maxPos.Y {
			continue
		}

		if (grid[Position{pos.X + 1, pos.Y}] == Scaffold &&
			grid[Position{pos.X - 1, pos.Y}] == Scaffold &&
			grid[Position{pos.X, pos.Y + 1}] == Scaffold &&
			grid[Position{pos.X, pos.Y - 1}] == Scaffold) {
			grid[pos] = Intersection

			calib += pos.X * pos.Y
		}
	}

	DrawGrid(grid, botPos, botDir)

	return fmt.Sprintf("%d", calib)
}

func DrawGrid(grid Grid, botPos Position, botDir star_15_1.Direction) {
	maxY := 0
	for pos, id := range grid {
		goterm.MoveCursor(pos.X + 1, pos.Y + 1)
		if botPos == pos {
			if id == Empty {
				goterm.Print("X")
				continue
			}

			switch botDir {
			case star_15_1.North:
				goterm.Print("^")
			case star_15_1.South:
				goterm.Print("v")
			case star_15_1.West:
				goterm.Print("<")
			case star_15_1.East:
				goterm.Print(">")
			}

			continue
		}

		switch id {
		case Scaffold:
			goterm.Print("#")
		case Empty:
			goterm.Print(".")
		case Intersection:
			goterm.Print("O")
		}

		if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	goterm.MoveCursor(1, maxY + 2)
	goterm.Flush()
}

func ReadAndRunProgram(file string) (chan<- int, <-chan int) {
	program := star_02_1.ReadProgramFromFile(file)

	input := make(chan int, 1)
	output := make(chan int)
	go func() {
		defer close(output)
		star_09_1.RunProgram(program, input, output)
	}()

	return input, output
}