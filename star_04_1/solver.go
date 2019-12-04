package star_04_1

import (
	"fmt"
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
		if AreDigitsValid(GetDigitsFromInt(i)) {
			validCount++
		}
	}

	return strconv.Itoa(validCount)
}

func GetDigitsFromInt (value int) []int {
	return GetDigitsFromString(strconv.Itoa(value))
}

func GetDigitsFromString (value string) []int {
	runes := []rune(value)
	digits := make([]int, len(runes))

	for i, r := range runes {
		digits[i], _ = strconv.Atoi(string(r))
	}

	return digits
}

func AreDigitsValid (digits []int) bool {
	areRaising := true
	hasPair := false

	for i, d := range digits {
		if i == 0 { continue }
		lastDigit := digits[i-1]
		if d < lastDigit { areRaising = false }
		if d == lastDigit { hasPair = true }
	}

	return areRaising && hasPair
}
