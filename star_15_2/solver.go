package star_15_2

import (
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_09_1"
	"github.com/buger/goterm"
	"strconv"
	"time"
)

type Solver struct {}

const File = "star_15_1/input"

type Position struct {
	X, Y int
}

type Bounds struct {
	Max, Min Position
}

type Direction int

const (
	North Direction = 1
	South Direction = 2
	West  Direction = 3
	East  Direction = 4
)

type StatusCode int

const (
	HitWall StatusCode = iota
	Moved
	Found
)

type GridID int

const (
	Unknown GridID = iota
	Origin
	Empty
	Wall
	O2S
)

type Grid map[Position]GridID

func (s *Solver) Solve() string {
	program := star_02_1.ReadProgramFromFile(File)

	input := make(chan int, 1)
	output := make(chan int)
	go func() {
		defer close(output)
		star_09_1.RunProgram(program, input, output)
	}()

	grid := make(Grid)

	var bounds Bounds
	var botPos Position
	botDir := North
	grid[botPos] = Origin
	var o2Pos Position
	resolvedMaze := false
	goterm.Clear()
	for {
		input <- int(botDir)
		status := StatusCode(<-output)

		switch status {
		case HitWall:
			wallPos := MovePosition(botPos, botDir)
			grid[wallPos] = Wall
			UpdateBounds(&bounds, wallPos)
			botDir = TurnRight(botDir)
		case Found:
			botPos = MovePosition(botPos, botDir)
			grid[botPos] = O2S
			o2Pos = botPos
			UpdateBounds(&bounds, botPos)
			botDir = TurnLeft(botDir)
		case Moved:
			botPos = MovePosition(botPos, botDir)
			if grid[botPos] == Origin {
				resolvedMaze = true
			}
			grid[botPos] = Empty
			UpdateBounds(&bounds, botPos)
			botDir = TurnLeft(botDir)
		}

		if resolvedMaze {
			break
		}
	}

	o2Fill := make(map[Position]bool)
	o2Fill[o2Pos] = true

	var minutes int
	for {

		for y, y2 := bounds.Max.Y, 0; y >= bounds.Min.Y; y-- {
			for x := bounds.Min.X; x <= bounds.Max.X; x++ {
				goterm.MoveCursor(x+1-bounds.Min.X, y2+1)
				current := Position{x, y}
				switch grid[current] {
				case Unknown: fallthrough
				case Wall:
					goterm.Print(goterm.Color("█", goterm.WHITE))
				case O2S:
					str := goterm.Color("O", goterm.BLACK)
					if o2Fill[current] {
						str = goterm.Background(str, goterm.RED)
					}
					goterm.Print(str)
				case Empty:
					str := goterm.Color(".", goterm.BLUE)
					if o2Fill[current] {
						str = goterm.Background(str, goterm.RED)
					}
					goterm.Print(str)
				}
			}
			y2++
			goterm.MoveCursor(1, bounds.Max.Y-bounds.Min.Y+2)
			goterm.Println("Minutes:", minutes)
			goterm.Flush()
		}

		foundEmptyNeighbor := false
		newFill := make(map[Position]bool)
		for pos, _ := range o2Fill {
			neighbors := []Position{MovePosition(pos, North), MovePosition(pos, South), MovePosition(pos, West), MovePosition(pos, East)}

			for _, neighbor := range neighbors {
				if grid[neighbor] == Empty && !o2Fill[neighbor] {
					newFill[neighbor] = true
					foundEmptyNeighbor = true
				}
			}
		}

		if foundEmptyNeighbor == false {
			break
		}

		minutes++

		for pos, _ := range newFill {
			o2Fill[pos] = true
		}

		time.Sleep(10 * time.Millisecond)
	}
	return strconv.Itoa(minutes)
}

func MovePosition(position Position, direction Direction) Position {
	switch direction {
	case North:
		position.Y++
	case South:
		position.Y--
	case West:
		position.X--
	case East:
		position.X++
	}
	return position
}

func UpdateBounds(bounds *Bounds, pos Position) {
	if pos.X < bounds.Min.X {
		bounds.Min.X = pos.X
	}
	if pos.Y < bounds.Min.Y {
		bounds.Min.Y = pos.Y
	}
	if pos.X > bounds.Max.X {
		bounds.Max.X = pos.X
	}
	if pos.Y > bounds.Max.Y {
		bounds.Max.Y = pos.Y
	}
}

func TurnRight(dir Direction) Direction {
	switch dir {
	case North:
		return East
	case South:
		return West
	case West:
		return North
	case East:
		return South
	}
	panic("unexpected direction")
}

func TurnLeft(dir Direction) Direction {
	switch dir {
	case North:
		return West
	case South:
		return East
	case West:
		return South
	case East:
		return North
	}
	panic("unexpected direction")
}
