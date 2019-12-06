package star_06_1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const file = "star_06_1/input"

type Orbit struct{
	body string
	satellite string
}

type Body struct{
	id string
	parent *Body
	children Children
}

type Children map[string]*Body

type Solver struct{}

func (s *Solver) Solve () string {
	return SolveForFile(file)
}

func SolveForFile(input string) string {
	defs := GetOrbitsFromFile(input)

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
		}
		body.children[sat.id] = sat
	}

	count := CountOrbits(0, bodies["COM"])

	return strconv.Itoa(count)
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

	defs := strings.Split(string(bytes), "\n")
	orbits := make([]Orbit, len(defs))
	for i, def := range defs {
		bodies := strings.Split(def, ")")
		orbits[i] = Orbit{bodies[0], bodies[1]}
	}
	return orbits
}