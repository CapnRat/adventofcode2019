package star_03_1

import (
	"reflect"
	"testing"
)

func TestFindClosestIntersectionDistance(t *testing.T) {
	type args struct {
		wires []Wire
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]Wire{
			BuildWireFromDefinition("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			BuildWireFromDefinition("U62,R66,U55,R34,D71,R55,D58,R83"),
			}}, 159,
		},
		{"Example 2", args{[]Wire{
			BuildWireFromDefinition("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			BuildWireFromDefinition("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
		}}, 135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindClosestIntersectionDistance(tt.args.wires); got != tt.want {
				t.Errorf("FindClosestIntersectionDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildWireFromDefinition(t *testing.T) {
	type args struct {
		def string
	}
	tests := []struct {
		name string
		args args
		want Wire
	}{
		{"Simple Wire" , args{"D3,R83,U83,L124"}, Wire{
			{Down, 3},
			{Right, 83},
			{Up, 83},
			{Left, 124},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildWireFromDefinition(tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildWireFromDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}