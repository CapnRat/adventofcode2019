package star_16_1

import "testing"

func TestSolveInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example", args{"80871224585914546619083218645595"}, "24176176"},
		{"Example", args{"19617804207202209144916044189917"}, "73745418"},
		{"Example", args{"69317163492948606335995924319873"}, "52432133"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveInput(tt.args.input); got != tt.want {
				t.Errorf("SolveInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveInputWithPatternForPhases(t *testing.T) {
	type args struct {
		input   string
		pattern []int
		phases  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example", args{"12345678", []int{0,1,0,-1}, 1}, "48226158"},
		{"Example", args{"12345678", []int{0,1,0,-1}, 2}, "34040438"},
		{"Example", args{"12345678", []int{0,1,0,-1}, 3}, "03415518"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveInputWithPatternForPhases(tt.args.input, tt.args.pattern, tt.args.phases); got != tt.want {
				t.Errorf("SolveInputWithPatternForPhases() = %v, want %v", got, tt.want)
			}
		})
	}
}