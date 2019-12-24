package star_10_1

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

const (
	CharEmpty    byte = '.'
	CharAsteroid byte = '#'
)

const file = "star_10_1/input"

type AsteroidField map[SpacePoint]bool

type SpacePoint struct{
	x int
	y int
}

type Asteroid struct{
	Point SpacePoint
	VisibleAsteroids int
}

type Solver struct{}

func (s *Solver) Solve() string {
	asteroid := SolveWithInputFile(file)
	return fmt.Sprintf("%d, %d (%d)", asteroid.Point.x, asteroid.Point.y, asteroid.VisibleAsteroids)
}

func SolveWithInputFile(path string) Asteroid {
	field := ParseAsteroidField(path)

	var best Asteroid
	for point := range *field {
		type radial struct{
			x int64
			y int64
		}
		radials := make(map[radial]int)

		for target := range *field {
			relativeX := target.x - point.x
			relativeY := target.y - point.y

			if relativeX == 0 && relativeY == 0 {
				continue
			}

			var rad radial
			if relativeX == 0 {
				rad = radial {0, int64(Sign(relativeY))}
			} else if relativeY == 0 {
				rad = radial {int64(Sign(relativeX)), 0}
			} else {
				a := big.NewInt(int64(relativeX))
				b := big.NewInt(int64(relativeY))
				absA := new(big.Int)
				absB := new(big.Int)
				absA.Abs(a)
				absB.Abs(b)
				var gcd big.Int
				gcd.GCD(nil, nil, absA, absB)
				rad = radial{a.Div(a, &gcd).Int64(), b.Div(b, &gcd).Int64()}
			}

			radials[rad]++
		}

		visible := len(radials)
		if best.VisibleAsteroids < visible {
			best.VisibleAsteroids = visible
			best.Point = point
		}
	}

	return best
}

func Sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
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
			case CharAsteroid:
				field[SpacePoint{x,y}] = true
			}
		}
		y++
	}

	return &field
}
