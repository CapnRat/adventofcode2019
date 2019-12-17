package star_10_1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	Empty    byte = '.'
	Asteroid byte = '#'
)

const file = "star_10_1/file"

type AsteroidField map[SpacePoint]bool

type SpacePoint struct{
	x int
	y int
}

type Solver struct{}

func (s *Solver) Solve() string {
	asteroid, count := SolveWithInputFile(file)
	return fmt.Sprintf("%d, %d (%d)", asteroid.x, asteroid.y, count)
}

func SolveWithInputFile(path string) (SpacePoint, int) {
	field := ParseAsteroidField(path)

	for asteroid := range *field {
		
	}

	return SpacePoint{}, 0
}

func ParseAsteroidField(path string) *AsteroidField {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	field := make(AsteroidField)
	y := 0
	for scanner.Scan() {
		for x, b := range scanner.Bytes() {
			switch b {
			case Asteroid:
				field[SpacePoint{x,y}] = true
			}
		}
	}

	return &field
}
