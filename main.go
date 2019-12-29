package main

import (
	"flag"
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_01_1"
	"github.com/CapnRat/adventofcode2019/star_01_2"
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_02_2"
	"github.com/CapnRat/adventofcode2019/star_03_1"
	"github.com/CapnRat/adventofcode2019/star_03_2"
	"github.com/CapnRat/adventofcode2019/star_04_1"
	"github.com/CapnRat/adventofcode2019/star_04_2"
	"github.com/CapnRat/adventofcode2019/star_05_1"
	"github.com/CapnRat/adventofcode2019/star_05_2"
	"github.com/CapnRat/adventofcode2019/star_06_1"
	"github.com/CapnRat/adventofcode2019/star_06_2"
	"github.com/CapnRat/adventofcode2019/star_07_1"
	"github.com/CapnRat/adventofcode2019/star_07_2"
	"github.com/CapnRat/adventofcode2019/star_08_1"
	"github.com/CapnRat/adventofcode2019/star_08_2"
	"github.com/CapnRat/adventofcode2019/star_09_1"
	"github.com/CapnRat/adventofcode2019/star_09_2"
	"github.com/CapnRat/adventofcode2019/star_10_1"
	"github.com/CapnRat/adventofcode2019/star_10_2"
	"github.com/CapnRat/adventofcode2019/star_11_1"
	"github.com/CapnRat/adventofcode2019/star_11_2"
	"github.com/CapnRat/adventofcode2019/star_12_1"
	"github.com/CapnRat/adventofcode2019/star_12_2"
	"github.com/CapnRat/adventofcode2019/star_13_1"
	"github.com/CapnRat/adventofcode2019/star_13_2"
	"github.com/CapnRat/adventofcode2019/star_14_1"
	"github.com/CapnRat/adventofcode2019/star_14_2"
	"github.com/CapnRat/adventofcode2019/star_15_1"
	"github.com/CapnRat/adventofcode2019/star_15_2"
	"github.com/CapnRat/adventofcode2019/star_16_1"
	"os"
	"time"
)

var solvers []Day

func registerSolvers() {
	solvers = append(solvers, Day{&star_01_1.Solver{}, &star_01_2.Solver{}})
	solvers = append(solvers, Day{&star_02_1.Solver{}, &star_02_2.Solver{}})
	solvers = append(solvers, Day{&star_03_1.Solver{}, &star_03_2.Solver{}})
	solvers = append(solvers, Day{&star_04_1.Solver{}, &star_04_2.Solver{}})
	solvers = append(solvers, Day{&star_05_1.Solver{}, &star_05_2.Solver{}})
	solvers = append(solvers, Day{&star_06_1.Solver{}, &star_06_2.Solver{}})
	solvers = append(solvers, Day{&star_07_1.Solver{}, &star_07_2.Solver{}})
	solvers = append(solvers, Day{&star_08_1.Solver{}, &star_08_2.Solver{}})
	solvers = append(solvers, Day{&star_09_1.Solver{}, &star_09_2.Solver{}})
	solvers = append(solvers, Day{&star_10_1.Solver{}, &star_10_2.Solver{}})
	solvers = append(solvers, Day{&star_11_1.Solver{}, &star_11_2.Solver{}})
	solvers = append(solvers, Day{&star_12_1.Solver{}, &star_12_2.Solver{}})
	solvers = append(solvers, Day{&star_13_1.Solver{}, &star_13_2.Solver{}})
	solvers = append(solvers, Day{&star_14_1.Solver{}, &star_14_2.Solver{}})
	solvers = append(solvers, Day{&star_15_1.Solver{}, &star_15_2.Solver{}})
	solvers = append(solvers, Day{&star_16_1.Solver{}, nil})
}

func main() {
	registerSolvers()

	var runLatest bool
	var runAll bool
	flag.BoolVar(&runLatest, "latest", false, "just run latest")
	flag.BoolVar(&runAll, "all", false, "run all solvers")
	flag.Parse()

	if runAll {
		for i, day := range solvers {
			RunSolver(i, 1)
			if day.star2 != nil {
				RunSolver(i, 2)
			}
		}
	} else if runLatest {
		RunSolver(getLatest())
	} else {
		RunSolver(getFromInput())
	}


}

func RunSolver(day int, star int) {
	fmt.Printf("Solving Day %d Star %d\n=====================\n", day+1, star)

	var solver Solver
	if star == 1 {
		solver = solvers[day].star1
	} else {
		solver = solvers[day].star2
	}

	start := time.Now()
	fmt.Println(solver.Solve())
	elapsed := time.Now().Sub(start)
	fmt.Printf("Finished after %d µs\n\n", elapsed/time.Microsecond)
}

func getLatest() (int, int) {
	day := len(solvers) - 1
	star := 1
	if solvers[day].star2 != nil {
		star = 2
	}
	return day, star
}

func getFromInput() (int, int) {
	listSolvers()
	fmt.Print("Enter Day: ")
	var day int
	n, err := fmt.Scanf("%d", &day)
	if err != nil || n != 1 {
		fmt.Println("error parsing day")
		os.Exit(1)
	}

	if len(solvers) < day || day < 1 {
		fmt.Println("error day out of range")
		os.Exit(1)
	}
	day--

	star := 1
	if solvers[day].star2 != nil {
		fmt.Print("Enter Star: ")
		n, err := fmt.Scanf("%d", &star)
		if err != nil || n != 1 {
			fmt.Println("error parsing star")
			os.Exit(1)
		}

		if n < 1 || n > 2 {
			fmt.Println("error star out of range")
			os.Exit(1)
		}
	}

	return day, star
}

func listSolvers() {
	fmt.Println("Solvers")
	for i, day := range solvers {
		stars := "1"
		if day.star2 != nil {
			stars += ", 2"
		}
		fmt.Printf("Day %d: (%s)\n", i+1, stars)
	}
}
