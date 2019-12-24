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
		want  Asteroid
	}{
		{"example0", args{"example0"}, Asteroid{SpacePoint{3, 4}, 8}},
		{"example1", args{"example1"}, Asteroid{SpacePoint{5, 8}, 33}},
		{"example2", args{"example2"}, Asteroid{SpacePoint{1, 2}, 35}},
		{"example3", args{"example3"}, Asteroid{SpacePoint{6, 3}, 41}},
		{"example4", args{"example4"}, Asteroid{SpacePoint{11, 13}, 210}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveWithInputFile(tt.args.file)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveWithInputFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}