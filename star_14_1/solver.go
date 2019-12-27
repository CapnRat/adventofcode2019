package star_14_1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Solver struct {}

type Element string

type Ingredients map[Element]int

type Recipe struct {
	in    Ingredients
	out   Element
	count int
}

type Batch struct {
	element Element
	count   int
}

type Recipes map[Element]Recipe

const File = "star_14_1/input"

func (s *Solver) Solve() string {
	return strconv.Itoa(SolveFromInput(File))
}

func SolveFromInput(path string) int {
	recipes := ReadRecipesFromInput(path)

	need := []Batch{{"FUEL", 1}}
	pool := make(map[Element]int)
	ore := 0

	for len(need) > 0 {
		batch := need[0]

		batch.count -= pool[batch.element]
		if batch.count <= 0 {
			pool[batch.element] = -batch.count
			need = need[1:]
			continue
		}
		pool[batch.element] = 0

		if batch.element == "ORE" {
			ore += batch.count
			need = need[1:]
			continue
		}

		recipe := recipes[batch.element]
		for element, count := range recipe.in {
			need = append(need, Batch{element, count})
		}
		batch.count -= recipe.count
		need[0] = batch
		if batch.count <= 0 {
			pool[batch.element] -= batch.count
			need = need[1:]
		}
	}

	return ore
}

func ReadRecipesFromInput(path string) Recipes {
	file, err := os.Open(path)
	if err != nil {
		panic (err)
	}

	recipes := make(Recipes)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		recipe := ParseLine(line)
		recipes[recipe.out] = recipe
	}

	return recipes
}

func ParseLine(line string) Recipe {
	reg := regexp.MustCompile(`(\d+) ([A-Z]+)`)
	matches := reg.FindAllStringSubmatch(line, -1)
	if len(matches) < 2 {
		panic("failed to parse line")
	}

	recipe := Recipe{ in: make(Ingredients) }

	for i, match := range matches {
		matchElement := Element(match[2])
		matchCount, err := strconv.Atoi(match[1])
		if err != nil {
			panic("failed to parse line")
		}
		if i == len(matches) - 1 {
			recipe.out = matchElement
			recipe.count = matchCount
		} else {
			recipe.in[matchElement] = matchCount
		}
	}

	return recipe
}

