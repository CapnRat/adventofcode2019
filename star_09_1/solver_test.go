package star_09_1

import (
	"github.com/CapnRat/adventofcode2019/star_02_1"
	"testing"
)

func TestSolveWithProgram(t *testing.T) {
	type args struct {
		program []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example", args{star_02_1.ParseProgram("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")}, "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"},
		{"Example", args{star_02_1.ParseProgram("104,1125899906842624,99")}, "1125899906842624"},
		{"Example", args{star_02_1.ParseProgram("1102,34915192,34915192,7,4,7,99,0")}, "1219070632396864"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveWithProgram(tt.args.program, nil); got != tt.want {
				t.Errorf("SolveWithProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}