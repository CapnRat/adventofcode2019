package star_06_2

import "testing"

func TestSolveForFile(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example", args{"testinput"}, "4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveForFile(tt.args.input); got != tt.want {
				t.Errorf("SolveForFile() = %v, want %v", got, tt.want)
			}
		})
	}
}