package star_03_2

import (
	"github.com/CapnRat/adventofcode2019/star_03_1"
	"testing"
)

func TestFindClosestIntersectionDistance(t *testing.T) {
	type args struct {
		wires []star_03_1.Wire
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]star_03_1.Wire{
			star_03_1.BuildWireFromDefinition("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			star_03_1.BuildWireFromDefinition("U62,R66,U55,R34,D71,R55,D58,R83"),
			}}, 610,
		},
		{"Example 2", args{[]star_03_1.Wire{
			star_03_1.BuildWireFromDefinition("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			star_03_1.BuildWireFromDefinition("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			}}, 410,
		},
		{"Example 3", args{[]star_03_1.Wire{
			star_03_1.BuildWireFromDefinition("R8,U5,L5,D3"),
			star_03_1.BuildWireFromDefinition("U7,R6,D4,L4"),
		}}, 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShortestIntersectionDistance(tt.args.wires); got != tt.want {
				t.Errorf("FindShortestIntersectionDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}