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
	In       Ingredients
	out      Element
	OutCount int
}

type Batch struct {
	Element Element
	Count   int
}

type Recipes map[Element]Recipe

const File = "star_14_1/input"

func (s *Solver) Solve() string {
	return strconv.Itoa(SolveFromInput(File))
}

func CalculateOreForFuel(recipes Recipes, fuel int) int {
	need := []Batch{{"FUEL", fuel}}
	pool := make(map[Element]int)
	ore := 0

	for len(need) > 0 {
		batch := need[0]

		batch.Count -= pool[batch.Element]
		if batch.Count <= 0 {
			pool[batch.Element] = -batch.Count
			need = need[1:]
			continue
		}
		pool[batch.Element] = 0

		if batch.Element == "ORE" {
			ore += batch.Count
			need = need[1:]
			continue
		}

		recipe := recipes[batch.Element]
		recipeCount := batch.Count / recipe.OutCount
		if batch.Count%recipe.OutCount > 0 {
			recipeCount++
		}
		for element, count := range recipe.In {
			need = append(need, Batch{Element: element, Count: count * recipeCount})
		}
		batch.Count -= recipe.OutCount * recipeCount
		need[0] = batch
		if batch.Count <= 0 {
			pool[batch.Element] -= batch.Count
			need = need[1:]
		}
	}

	return ore
}

func SolveFromInput(path string) int {
	recipes := ReadRecipesFromInput(path)
	return CalculateOreForFuel(recipes, 1)
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

	recipe := Recipe{ In: make(Ingredients) }

	for i, match := range matches {
		matchElement := Element(match[2])
		matchCount, err := strconv.Atoi(match[1])
		if err != nil {
			panic("failed to parse line")
		}
		if i == len(matches) - 1 {
			recipe.out = matchElement
			recipe.OutCount = matchCount
		} else {
			recipe.In[matchElement] = matchCount
		}
	}

	return recipe
}

