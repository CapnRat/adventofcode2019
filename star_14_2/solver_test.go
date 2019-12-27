package star_14_2

import "testing"

func TestSolveFromInput(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example3", args{"../star_14_1/example3"}, 82892753},
		{"Example4", args{"../star_14_1/example4"}, 5586022},
		{"Example5", args{"../star_14_1/example5"}, 460664},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveFromInput(tt.args.path); got != tt.want {
				t.Errorf("SolveFromInput() = %v, want %v", got, tt.want)
			}
		})
	}
}