package star_03_1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const File = "star_03_1/input"

const (
	Up = iota
	Down
	Right
	Left
)
type direction int

type step struct{
	Direction direction
	Length    int
}

type Wire []step

type Coordinate struct{
	X int
	Y int
}

type Solver struct{}

func (s *Solver) Solve () string {
	file, err := os.Open(File)
	if err != nil {
		fmt.Println("failed to load file ", err)
		os.Exit(1)
	}
	defer file.Close()

	wires := make([]Wire, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wires = append(wires, BuildWireFromDefinition(scanner.Text()))
	}

	result := FindClosestIntersectionDistance(wires)

	return strconv.Itoa(result)
}

func BuildWireFromDefinition (def string) Wire {
	var wire Wire

	for _, defstep := range strings.Split(def, ",") {
		var step step
		switch defstep[0] {
		case 'U':
			step.Direction = Up
		case 'D':
			step.Direction = Down
		case 'R':
			step.Direction = Right
		case 'L':
			step.Direction = Left
		default:
			fmt.Println("error parsing def direction")
			os.Exit(1)
		}
		length, err := strconv.Atoi(string([]rune(defstep[1:])))
		if err != nil {
			fmt.Println("error parsing def length")
			os.Exit(1)
		}
		step.Length = length

		wire = append(wire, step)
	}

	return wire
}

func FindClosestIntersectionDistance (wires []Wire) int {
	visited := make(map[Coordinate]int)
	distance := math.MaxInt64

	for i, wire := range wires {
		cursor := Coordinate{0, 0}
		for _, step := range wire {
			for j := 0; j < step.Length; j++ {
				switch step.Direction {
				case Up:
					cursor.Y++
				case Down:
					cursor.Y--
				case Right:
					cursor.X++
				case Left:
					cursor.X--
				}

				if visitor, found := visited[cursor]; found && visitor != i {
					// already visited by another wire
					cursorDistance := Abs(cursor.X) + Abs(cursor.Y)
					if cursorDistance < distance {
						distance = cursorDistance
					}
				} else {
					visited[cursor] = i
				}
			}
		}
	}

	return distance
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
