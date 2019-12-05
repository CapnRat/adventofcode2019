package star_05_2

import (
	"reflect"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	type args struct {
		instruction int
	}
	tests := []struct {
		name       string
		args       args
		wantOpCode OpCode
		wantModes  []ParamMode
		wantLength int
	}{
		{"Null", args{0}, OpNul, nil, 1},
		{"AddWithPositionModes", args{1}, OpAdd, []ParamMode{PositionMode, PositionMode}, 4},
		{"AddWithImmediateModes", args{1101}, OpAdd, []ParamMode{ImmediateMode, ImmediateMode}, 4},
		{"AddWithMixedModes", args{1001}, OpAdd, []ParamMode{PositionMode, ImmediateMode}, 4},
		{"MulWithPositionModes", args{2}, OpMul, []ParamMode{PositionMode, PositionMode}, 4},
		{"MulWithImmediateModes", args{1102}, OpMul, []ParamMode{ImmediateMode, ImmediateMode}, 4},
		{"MulWithMixedModes", args{1002}, OpMul, []ParamMode{PositionMode, ImmediateMode}, 4},
		{"In", args{3}, OpIn, nil, 2},
		{"OutWithPositionMode", args{4}, OpOut, []ParamMode{PositionMode}, 2},
		{"OutWithImmediateMode", args{104}, OpOut, []ParamMode{ImmediateMode}, 2},
		{"JumpIfTrueWithPositionModes", args{5}, OpJmpT, []ParamMode{PositionMode, PositionMode}, 0},
		{"JumpIfTrueWithImmediateModes", args{1105}, OpJmpT, []ParamMode{ImmediateMode, ImmediateMode}, 0},
		{"JumpIfTrueWithMixedModes", args{1005}, OpJmpT, []ParamMode{PositionMode, ImmediateMode}, 0},
		{"JumpIfFalseWithPositionModes", args{6}, OpJmpF, []ParamMode{PositionMode, PositionMode}, 0},
		{"JumpIfFalseWithImmediateModes", args{1106}, OpJmpF, []ParamMode{ImmediateMode, ImmediateMode}, 0},
		{"JumpIfFalseWithMixedModes", args{1006}, OpJmpF, []ParamMode{PositionMode, ImmediateMode}, 0},
		{"LessThanWithPositionModes", args{7}, OpLess, []ParamMode{PositionMode, PositionMode}, 4},
		{"LessThanWithImmediateModes", args{1107}, OpLess, []ParamMode{ImmediateMode, ImmediateMode}, 4},
		{"LessThanWithMixedModes", args{1007}, OpLess, []ParamMode{PositionMode, ImmediateMode}, 4},
		{"EqualsWithPositionModes", args{8}, OpEql, []ParamMode{PositionMode, PositionMode}, 4},
		{"EqualsWithImmediateModes", args{1108}, OpEql, []ParamMode{ImmediateMode, ImmediateMode}, 4},
		{"EqualsWithMixedModes", args{1008}, OpEql, []ParamMode{PositionMode, ImmediateMode}, 4},
		{"Halt", args{99}, OpHalt, nil, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpCode, gotModes, gotLength := ParseInstruction(tt.args.instruction)
			if gotOpCode != tt.wantOpCode {
				t.Errorf("ParseInstruction() gotOpcode = %v, want %v", gotOpCode, tt.wantOpCode)
			}
			if !reflect.DeepEqual(gotModes, tt.wantModes) {
				t.Errorf("ParseInstruction() gotModes = %v, want %v", gotModes, tt.wantModes)
			}
			if gotLength != tt.wantLength {
				t.Errorf("ParseInstruction() gotLength = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}