package star_12_1

import (
	"bufio"
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_03_1"
	"os"
	"regexp"
	"strconv"
)

type Solver struct {}

type Moon struct {
	name string
	x, y, z int
	velx, vely, velz int
}

var MoonNames = []string{"Io", "Europa", "Ganymede", "Callisto"}

var Moons []Moon

const File = "star_12_1/input"

func (s *Solver) Solve() string {
	return strconv.Itoa(SolveForInput(File,1000))
}

func SolveForInput(path string, steps int) int {
	moons := GetMoonsFromInput(path)
	SimulateForSteps(moons, steps)
	return CalculateEnergy(moons)
}

func CalculateEnergy(moons []Moon) (energy int) {
	for _, moon := range moons {
		pot := star_03_1.Abs(moon.x) + star_03_1.Abs(moon.y) + star_03_1.Abs(moon.z)
		kin := star_03_1.Abs(moon.velx) + star_03_1.Abs(moon.vely) + star_03_1.Abs(moon.velz)
		energy += pot * kin
	}
	return
}

func SimulateForSteps(moons []Moon, steps int) {
	for i := 0; i < steps; i++ {
		SimulateMoons(moons)
	}
}

func SimulateMoons(moons []Moon) {
	// Update Velocities
	for i, _ := range moons {
		moon := &moons[i]
		for j := i + 1; j < len(moons); j++ {
			other := &moons[j]
			if moon.x < other.x {
				moon.velx++
				other.velx--
			}
			if moon.x > other.x {
				moon.velx--
				other.velx++
			}
			if moon.y < other.y {
				moon.vely++
				other.vely--
			}
			if moon.y > other.y {
				moon.vely--
				other.vely++
			}
			if moon.z < other.z {
				moon.velz++
				other.velz--
			}
			if moon.z > other.z {
				moon.velz--
				other.velz++
			}
		}
	}

	// Update Positions
	for i, _ := range moons {
		moon := &moons[i]
		moon.x += moon.velx
		moon.y += moon.vely
		moon.z += moon.velz
	}
}

func GetMoonsFromInput(path string) []Moon {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("failed to load file ", err)
		os.Exit(1)
	}
	defer file.Close()

	var moons []Moon

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		exp := regexp.MustCompile(`x=(-?\d*).*y=(-?\d*).*z=(-?\d*)`)
		matches := exp.FindStringSubmatch(line)
		if len(matches) != 4 {
			panic("unable to parse line")
		}
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		z, _ := strconv.Atoi(matches[3])
		moons = append(moons, Moon{MoonNames[len(moons)], x, y, z, 0, 0, 0})
	}

	return moons
}
