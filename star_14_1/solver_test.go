package star_14_1

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Recipe
	}{
		{"Example", args{"10 ORE => 10 A"}, Recipe{
			In:    Ingredients {
				"ORE": 10,
			},
			out:      "A",
			OutCount: 10,
		}},
		{"Example", args{"7 A, 1 B => 1 C"}, Recipe{
			In:    Ingredients {
				"A": 7,
				"B": 1,
			},
			out:      "C",
			OutCount: 1,
		}},
		{"Example", args{"2 AB, 3 BC, 4 CA => 1 FUEL"}, Recipe{
			In:    Ingredients {
				"AB": 2,
				"BC": 3,
				"CA": 4,
			},
			out:      "FUEL",
			OutCount: 1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveFromInput(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example1", args{"example1"}, 31},
		{"Example2", args{"example2"}, 165},
		{"Example3", args{"example3"}, 13312},
		{"Example4", args{"example4"}, 180697},
		{"Example5", args{"example5"}, 2210736},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveFromInput(tt.args.path); got != tt.want {
				t.Errorf("SolveFromInput() = %v, want %v", got, tt.want)
			}
		})
	}
}