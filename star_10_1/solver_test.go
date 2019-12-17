package star_10_1

import (
	"reflect"
	"testing"
)

func TestSolveWithInputFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name  string
		args  args
		want  SpacePoint
		want1 int
	}{
		{"example0", args{"star_10_1/example0"}, SpacePoint{3, 4}, 8},
		{"example1", args{"star_10_1/example1"}, SpacePoint{5, 8}, 33},
		{"example2", args{"star_10_1/example2"}, SpacePoint{1, 2}, 35},
		{"example3", args{"star_10_1/example3"}, SpacePoint{6, 3}, 41},
		{"example4", args{"star_10_1/example4"}, SpacePoint{11, 13}, 210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SolveWithInputFile(tt.args.file)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveWithInputFile() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SolveWithInputFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}