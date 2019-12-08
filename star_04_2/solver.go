package star_04_2

import (
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_04_1"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const file = "star_04_1/input"

type Solver struct{}

func (s *Solver) Solve () string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	bounds := strings.Split(string(bytes), "-")
	if len(bounds) != 2 {
		fmt.Println("error parsing file")
		os.Exit(1)
	}

	min, _ := strconv.Atoi(bounds[0])
	max, _ := strconv.Atoi(bounds[1])

	validCount := 0
	for i := min; i <= max; i++ {
		if AreDigitsValid(star_04_1.GetDigitsFromInt(i)) {
			validCount++
		}
	}

	return strconv.Itoa(validCount)
}

func AreDigitsValid (digits []int) bool {
	groupCounts := make(map[int]int)

	for i, d := range digits {
		if i == 0 {
			groupCounts[d] = 1
			continue
		}
		lastDigit := digits[i-1]
		if d < lastDigit { return false }
		if _, ok := groupCounts[d]; ok {
			groupCounts[d]++
		} else {
			groupCounts[d] = 1
		}
	}

	hasPair := false
	for n := 0; n < 10; n++ {
		if count, ok := groupCounts[n]; ok && count == 2 {
			hasPair = true
		}
	}

	return hasPair
}
