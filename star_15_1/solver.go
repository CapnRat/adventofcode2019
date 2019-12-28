package star_15_1

import (
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_09_1"
	"github.com/buger/goterm"
	"strconv"
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
	path := make(map[Position]bool)

	var bounds Bounds
	var botPos Position
	botDir := North
	grid[botPos] = Origin
	found := false
	goterm.Clear()
	for {
		input <- int(botDir)
		status := StatusCode(<- output)

		switch status {
		case HitWall:
			wallPos := MovePosition(botPos, botDir)
			grid[wallPos] = Wall
			UpdateBounds(&bounds, wallPos)
			botDir = TurnRight(botDir)
		case Found:
			prevPos := botPos
			botPos = MovePosition(botPos, botDir)
			grid[botPos] = O2S
			if path[botPos] {
				delete(path, prevPos)
			} else {
				path[botPos] = true
			}
			UpdateBounds(&bounds, botPos)
			botDir = TurnLeft(botDir)
			found = true
		case Moved:
			prevPos := botPos
			botPos = MovePosition(botPos, botDir)
			if grid[botPos] == Origin {
				goterm.MoveCursor(1, bounds.Max.Y - bounds.Min.Y + 2)
				goterm.Flush()
				return strconv.Itoa(len(path))
			}
			grid[botPos] = Empty
			if !found {
				if path[botPos] {
					delete(path, prevPos)
				} else {
					path[botPos] = true
				}
			}
			UpdateBounds(&bounds, botPos)
			botDir = TurnLeft(botDir)
		}
		for y, y2 := bounds.Max.Y, 0; y >= bounds.Min.Y; y-- {
			for x := bounds.Min.X; x <= bounds.Max.X; x++ {
				goterm.MoveCursor(x+1-bounds.Min.X, y2+1)
				if botPos.X == x && botPos.Y == y {
					switch botDir {
					case North:
						goterm.Print("^")
					case South:
						goterm.Print("v")
					case West:
						goterm.Print("<")
					case East:
						goterm.Print(">")
					}
				} else {
					current := Position{x, y}
					switch grid[current] {
					case Wall:
						goterm.Print(goterm.Color("█", goterm.WHITE))
					case O2S:
						str := goterm.Color("O", goterm.RED)
						if path[current] {
							str = goterm.Background(str, goterm.CYAN)
						}
						goterm.Print(str)
					case Empty:
						str := goterm.Color(".", goterm.BLUE)
						if path[current] {
							str = goterm.Background(str, goterm.CYAN)
						}
						goterm.Print(str)
					case Unknown:
						goterm.Print(" ")
					case Origin:
						str := goterm.Color("▲", goterm.GREEN)
						if path[current] {
							str = goterm.Background(str, goterm.CYAN)
						}
						goterm.Print(str)
					}
				}
			}
			y2++
		}
		goterm.MoveCursor(1, bounds.Max.Y - bounds.Min.Y + 2)
		goterm.Flush()
		//time.Sleep(10*time.Millisecond)
	}
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
