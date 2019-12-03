package star_03_2

import (
	"bufio"
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_03_1"
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
	}

	visited := make(map[star_03_1.Coordinate]Visit)
	distance := math.MaxInt64

	for i, wire := range wires {
		cursor := star_03_1.Coordinate{0, 0}
		stepCount := 0
		for _, step := range wire {
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
				} else {
					visited[cursor] = Visit{stepCount, i}
				}
			}
		}
	}

	return distance
}
