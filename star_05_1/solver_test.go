package star_05_1

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