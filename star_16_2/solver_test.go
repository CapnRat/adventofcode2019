package star_16_2

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
		{"Example", args{"03036732577212944063491565474664"}, "84462026"},
		{"Example", args{"02935109699940807407585447034323"}, "78725270"},
		{"Example", args{"03081770884921959731165446850517"}, "53553731"},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveInput([]byte(tt.args.input)); got != tt.want {
				t.Errorf("SolveInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
