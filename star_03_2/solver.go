package star_03_2

import (
	"bufio"
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_03_1"
	"github.com/gookit/color"
	"math"
	"os"
	"strconv"
)

type Solver struct{}

func (s *Solver) Solve () string {
	file, err := os.Open(star_03_1.File)
	if err != nil {
		fmt.Println("failed to load file ", err)
		os.Exit(1)
	}
	defer file.Close()

	wires := make([]star_03_1.Wire, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wires = append(wires, star_03_1.BuildWireFromDefinition(scanner.Text()))
	}

	result := FindShortestIntersectionDistance(wires)

	return strconv.Itoa(result)
}

func FindShortestIntersectionDistance(wires []star_03_1.Wire) int {
	type Visit struct{
		numSteps int
		visitor int
		direction star_03_1.Direction
		corner bool
		intersection bool
	}

	type Bounds struct{
		maxX int
		minX int
		maxY int
		minY int
	}

	visited := make(map[star_03_1.Coordinate]Visit)
	distance := math.MaxInt64

	const buffer = 1
	bounds := Bounds{buffer, -buffer, buffer, -buffer}

	for i, wire := range wires {
		cursor := star_03_1.Coordinate{0, 0}
		stepCount := 0
		for k, step := range wire {
			for j := 0; j < step.Length; j++ {
				stepCount++

				switch step.Direction {
				case star_03_1.Up:
					cursor.Y++
				case star_03_1.Down:
					cursor.Y--
				case star_03_1.Right:
					cursor.X++
				case star_03_1.Left:
					cursor.X--
				}

				if visit, found := visited[cursor]; found && visit.visitor != i {
					// already visited by another wire
					cursorDistance := visit.numSteps + stepCount
					if cursorDistance < distance {
						distance = cursorDistance
					}
					visit.intersection = true
					visited[cursor] = visit
				} else {
					visited[cursor] = Visit{stepCount, i, step.Direction, j == step.Length - 1 && k < len(wire) - 1, false}
				}

				if cursor.X - buffer < bounds.minX { bounds.minX = cursor.X - buffer }
				if cursor.X + buffer > bounds.maxX { bounds.maxX = cursor.X + buffer }
				if cursor.Y - buffer < bounds.minY { bounds.minY = cursor.Y - buffer }
				if cursor.Y + buffer > bounds.maxY { bounds.maxY = cursor.Y + buffer }
			}
		}
	}

	for j := bounds.maxY; j >= bounds.minY; j-- {
		for i := bounds.minX; i <= bounds.maxX; i++ {
			outColor := color.Gray
			char := "."
			if visit, found := visited[star_03_1.Coordinate{X: i, Y: j}]; found {
				if visit.corner {
					char = "+"
					outColor = color.White
				} else if visit.intersection {
					char = "X"
					outColor = color.White
				} else if visit.direction == star_03_1.Up || visit.direction == star_03_1.Down {
					char = "|"
					if visit.visitor == 0 {
						outColor = color.Red
					} else {
						outColor = color.Green
					}
				} else {
					char = "-"
					if visit.visitor == 0 {
						outColor = color.Red
					} else {
						outColor = color.Green
					}
				}
			}
			if i == 0 && j == 0 {
				char = "o"
				outColor = color.White
			}
			outColor.Print(char)
		}
		fmt.Print("\n")
	}

	return distance
}
