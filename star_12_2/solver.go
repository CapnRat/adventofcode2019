package star_12_2

import (
	"github.com/CapnRat/adventofcode2019/star_12_1"
	"math/big"
	"strconv"
)

type Solver struct {}

const File = "star_12_1/input"

func (s *Solver) Solve() string {
	return strconv.Itoa(SolveForInput(File))
}

func SolveForInput(path string) int {
	moons := star_12_1.GetMoonsFromInput(path)

	periods := [3]int{}
	for i := 1;;i++ {
		SimulateMoons(moons)

		if periods[0] == 0 && moons[0].Velx == 0 && moons[1].Velx == 0 && moons[2].Velx == 0 && moons[3].Velx == 0 {
			periods[0] = i * 2
		}
		if periods[1] == 0 && moons[0].Vely == 0 && moons[1].Vely == 0 && moons[2].Vely == 0 && moons[3].Vely == 0 {
			periods[1] = i * 2
		}
		if periods[2] == 0 && moons[0].Velz == 0 && moons[1].Velz == 0 && moons[2].Velz == 0 && moons[3].Velz == 0 {
			periods[2] = i * 2
		}

		if periods[0] != 0 && periods[1] != 0 && periods[2] != 0 {
			break
		}
	}

	return LCM(LCM(periods[0], periods[1]), periods[2])
}

func GCD(a, b int) int {
	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))
	var gcd big.Int
	gcd.GCD(nil, nil, bigA, bigB)
	return int(gcd.Int64())
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func SimulateMoons(moons []star_12_1.Moon) {
	for i, _ := range moons {
		moon := &moons[i]
		for j := i + 1; j < len(moons); j++ {
			other := &moons[j]

			// Update Velocities
			if moon.X < other.X {
				moon.Velx++
				other.Velx--
			}
			if moon.X > other.X {
				moon.Velx--
				other.Velx++
			}
			if moon.Y < other.Y {
				moon.Vely++
				other.Vely--
			}
			if moon.Y > other.Y {
				moon.Vely--
				other.Vely++
			}
			if moon.Z < other.Z {
				moon.Velz++
				other.Velz--
			}
			if moon.Z > other.Z {
				moon.Velz--
				other.Velz++
			}
		}

		// Update position
		moon.X += moon.Velx
		moon.Y += moon.Vely
		moon.Z += moon.Velz
	}
}