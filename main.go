package main

import (
	"flag"
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_01_1"
	"github.com/CapnRat/adventofcode2019/star_01_2"
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"github.com/CapnRat/adventofcode2019/star_02_2"
	"github.com/CapnRat/adventofcode2019/star_03_1"
	"os"
)

var solvers []Day

func registerSolvers() {
	solvers = append(solvers, Day{&star_01_1.Solver{}, &star_01_2.Solver{}})
	solvers = append(solvers, Day{&star_02_1.Solver{}, &star_02_2.Solver{}})
	solvers = append(solvers, Day{&star_03_1.Solver{}, nil})
}

func main() {
	registerSolvers()

	var runLatest bool
	flag.BoolVar(&runLatest, "latest", false, "just run latest")
	flag.Parse()

	var day, star int
	if runLatest {
		day, star = getLatest()
	} else {
		day, star = getFromInput()
	}

	fmt.Printf("Solving Day %d Star %d\n=====================\n", day+1, star)
	var solver Solver
	if star == 1 {
		solver = solvers[day].star1
	} else {
		solver = solvers[day].star2
	}

	fmt.Println(solver.Solve())
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
