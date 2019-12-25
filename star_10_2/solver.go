package star_10_2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"sort"
)

const (
	CharEmpty    byte = '.'
	CharAsteroid byte = '#'
)

const file = "star_10_1/input"

type AsteroidField []SpacePoint

type SpacePoint struct{
	x int
	y int
}

type Asteroid struct{
	Point SpacePoint
	VisibleAsteroids int
}

type Radial struct{
	X int64
	Y int64
}

type Solver struct{}

func (s *Solver) Solve() string {
	asteroid := SolveWithInputFile(file)
	return fmt.Sprintf("%d, %d (%d)", asteroid.x, asteroid.y, EncodeSpacePoint(asteroid))
}

func EncodeSpacePoint(point SpacePoint) int {
	return point.x * 100 + point.y
}

func GetRadial(source SpacePoint, target SpacePoint) Radial {
	relativeX := target.x - source.x
	relativeY := target.y - source.y

	if relativeX == 0 && relativeY == 0 {
		return Radial{}
	}

	var rad Radial
	if relativeX == 0 {
		rad = Radial{0, int64(Sign(relativeY))}
	} else if relativeY == 0 {
		rad = Radial{int64(Sign(relativeX)), 0}
	} else {
		a := big.NewInt(int64(relativeX))
		b := big.NewInt(int64(relativeY))
		absA := new(big.Int)
		absB := new(big.Int)
		absA.Abs(a)
		absB.Abs(b)
		var gcd big.Int
		gcd.GCD(nil, nil, absA, absB)
		rad = Radial{a.Div(a, &gcd).Int64(), b.Div(b, &gcd).Int64()}
	}

	return rad
}

func FindDeathAsteroid(field AsteroidField) Asteroid {
	var best Asteroid
	for _, point := range field {
		radials := make(map[Radial]int)

		for _, target := range field {
			if point != target {
				radials[GetRadial(point, target)]++
			}
		}

		visible := len(radials)
		if best.VisibleAsteroids < visible {
			best.VisibleAsteroids = visible
			best.Point = point
		}
	}

	return best
}

func SolveWithInputFile(path string) SpacePoint {
	field := ParseAsteroidField(path)
	deathAsteroid := FindDeathAsteroid(field)

	field = SortField(field, deathAsteroid)

	return field[199]
}

func SortField (field AsteroidField, deathAsteroid Asteroid) AsteroidField {
	type radialList struct {
		radial Radial
		asteroids []SpacePoint
		frontIndex int
	}
	var radialLists []radialList
	seenRadials := make(map[Radial]int)
	for _, asteroid := range field {
		rad := GetRadial(deathAsteroid.Point, asteroid)
		if rad.X == 0 && rad.Y == 0 {
			continue
		}
		if _, ok := seenRadials[rad]; !ok {
			seenRadials[rad] = len(radialLists)
			radialLists = append(radialLists, radialList{
				radial:    rad,
				asteroids: nil,
				frontIndex: 0,
			})
		}
		radialLists[seenRadials[rad]].asteroids = append(radialLists[seenRadials[rad]].asteroids, asteroid)
	}

	sort.SliceStable(radialLists, func(i, j int) bool {
		EmptyRadial := Radial{0,0}
		if radialLists[i].radial == EmptyRadial {
			return false
		}
		if radialLists[j].radial == EmptyRadial {
			return true
		}
		return AngleBetweenRadials(Radial{0,-1}, radialLists[i].radial) < AngleBetweenRadials(Radial{0,-1}, radialLists[j].radial)
	})

	for _, radials := range radialLists {
		sort.SliceStable(radials.asteroids, func(i, j int) bool {
			return (radials.asteroids[i].x + radials.asteroids[i].y) < (radials.asteroids[j].x + radials.asteroids[j].y)
		})
	}

	sortedField := make([]SpacePoint, len(field) - 1)

	listIndex := 0
	for i := 0; i < len(sortedField); {
		list := &radialLists[listIndex]
		listIndex = (listIndex + 1) % len(radialLists)
		if list.frontIndex == len(list.asteroids) {
			continue
		}
		sortedField[i] = list.asteroids[list.frontIndex]
		list.frontIndex++
		i++
	}

	return sortedField
}

func AngleBetweenRadials(i Radial, j Radial) float64 {
	angle := math.Atan2(float64(i.X), float64(i.Y)) - math.Atan2(float64(j.X), float64(j.Y))
	if angle < 0 {
		angle += math.Pi * 2
	}
	return angle
}

func Sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

func ParseAsteroidField(path string) AsteroidField {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var field AsteroidField
	y := 0
	for scanner.Scan() {
		for x, b := range scanner.Bytes() {
			switch b {
			case CharAsteroid:
				field = append(field, SpacePoint{x,y})
			}
		}
		y++
	}

	return field
}
