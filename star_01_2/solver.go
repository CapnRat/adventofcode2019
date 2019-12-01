package star_01_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const file = "star_01_1/input"

type Solver struct{}

func (s *Solver) Solve() string {
	file, err := os.Open(file)
	if err != nil {
		fmt.Println("failed to load file ", err)
		os.Exit(1)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var mass int
		_, err := fmt.Sscanf(line, "%d", &mass)
		if err != nil {
			fmt.Println("failed to parse line ", line)
			os.Exit(1)
		}

		fuel := s.CalculateFuel(mass)

		total += fuel
		fmt.Printf("Mass: %d Fuel: %d Total: %d\n", mass, fuel, total)
	}
	return strconv.Itoa(total)
}

func (s *Solver) CalculateFuel(mass int) int {
	fuel := int(float64(mass)/3.0) - 2
	if fuel < 0 {
		fuel = 0
	} else {
		fuel += s.CalculateFuel(fuel)
	}
	return fuel
}
