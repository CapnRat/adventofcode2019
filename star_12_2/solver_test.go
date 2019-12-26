package star_12_2

import "testing"

func TestSolveForInput(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Fast", args{"../star_12_1/example"}, 2772},
		{"Slow", args{"example"}, 4686774924},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveForInput(tt.args.path); got != tt.want {
				t.Errorf("SolveForInput() = %v, want %v", got, tt.want)
			}
		})
	}
}