package star_02_1

import (
	"reflect"
	"testing"
)

func TestSolver_RunProgram(t *testing.T) {
	type args struct {
		program []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1,0,0,0,99", args{[]int{1,0,0,0,99}}, []int{2,0,0,0,99}},
		{"2,3,0,3,99", args{[]int{2,3,0,3,99}}, []int{2,3,0,6,99}},
		{"2,4,4,5,99,0", args{[]int{2,4,4,5,99,0}}, []int{2,4,4,5,99,9801}},
		{"1,1,1,4,99,5,6,0,99", args{[]int{1,1,1,4,99,5,6,0,99}}, []int{30,1,1,4,2,5,6,0,99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{}
			if got := s.RunProgram(tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}