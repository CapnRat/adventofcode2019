package star_06_2

import (
	"github.com/CapnRat/adventofcode2019/star_06_1"
	"strconv"
)

type Path []*star_06_1.Body

type Solver struct {}

func (s *Solver) Solve() string {
	return SolveForFile(star_06_1.File)
}

func SolveForFile(input string) string {
	defs := star_06_1.GetOrbitsFromFile(input)

	_, bodies := star_06_1.BuildOrbitTree(defs)

	count := CountStepsBetweenBodies(bodies["YOU"], bodies["SAN"])

	return strconv.Itoa(count)
}

func CountStepsBetweenBodies(src *star_06_1.Body, dst *star_06_1.Body) int {
	pathToSrc := GetPathToBody(src)
	pathToDst := GetPathToBody(dst)

	for i := 0; true; i++ {
		if pathToSrc[i] != pathToDst[i] {
			return len(pathToSrc) + len(pathToDst) - i * 2
		}
	}

	panic("fail")
}

func GetPathToBody(body *star_06_1.Body) Path {
	var pathBack Path
	pathBack = BuildPathBack(pathBack, body.Parent)
	var path Path
	for i := len(pathBack) - 1; i >= 0; i-- {
		path = append(path, pathBack[i])
	}
	return path
}

func BuildPathBack(path Path, body *star_06_1.Body) Path {
	path = append(path, body)
	if body.Parent != nil {
		path = BuildPathBack(path, body.Parent)
	}
	return path
}