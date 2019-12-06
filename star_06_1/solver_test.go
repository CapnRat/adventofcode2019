package star_06_1

import "testing"

func TestSolveForFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example", args{"testinput"}, "42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveForFile(tt.args.file); got != tt.want {
				t.Errorf("SolveForFile() = %v, want %v", got, tt.want)
			}
		})
	}
}