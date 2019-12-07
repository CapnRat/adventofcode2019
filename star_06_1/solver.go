package star_06_1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const File = "star_06_1/input"

type Orbit struct{
	body string
	satellite string
}

type Body struct{
	id       string
	Parent   *Body
	children Children
}

type Children map[string]*Body

type Solver struct{}

func (s *Solver) Solve () string {
	return SolveForFile(File)
}

func SolveForFile(input string) string {
	defs := GetOrbitsFromFile(input)

	com, _ := BuildOrbitTree(defs)

	count := CountOrbits(0, com)

	return strconv.Itoa(count)
}

func BuildOrbitTree (defs []Orbit) (*Body, map[string]*Body) {
	bodies := make(map[string]*Body)
	for _, def := range defs {
		body, ok := bodies[def.body]
		if !ok {
			body = &Body{def.body, nil, make(Children)}
			bodies[body.id] = body
		}
		sat, ok := bodies[def.satellite]
		if !ok {
			sat = &Body{def.satellite, body, make(Children)}
			bodies[sat.id] = sat
		} else {
			sat.Parent = body
		}
		body.children[sat.id] = sat
	}

	return bodies["COM"], bodies
}

func CountOrbits(depth int, body *Body) int {
	count := depth
	for _, sat := range body.children {
		count += CountOrbits(depth + 1, sat)
	}
	return count
}

func GetOrbitsFromFile(input string) []Orbit {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("error reading file", err)
		os.Exit(1)
	}

	defs := strings.Split(strings.Replace(string(bytes), "\r\n", "\n", -1), "\n")
	orbits := make([]Orbit, len(defs))
	for i, def := range defs {
		bodies := strings.Split(def, ")")
		orbits[i] = Orbit{bodies[0], bodies[1]}
	}
	return orbits
}