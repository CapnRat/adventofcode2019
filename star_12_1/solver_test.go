package star_12_1

import "testing"

func TestSolveForInput(t *testing.T) {
	type args struct {
		path string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{"example", 10}, 179},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveForInput(tt.args.path, tt.args.steps); got != tt.want {
				t.Errorf("SolveForInput() = %v, want %v", got, tt.want)
			}
		})
	}
}