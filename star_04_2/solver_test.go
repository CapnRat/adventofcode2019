package star_04_2

import (
	"github.com/CapnRat/adventofcode2019/star_04_1"
	"testing"
)

func TestAreDigitsValid(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid 1", args{star_04_1.GetDigitsFromString("122345")}, true},
		{"Valid 2", args{star_04_1.GetDigitsFromString("112233")}, true},
		{"Valid 3", args{star_04_1.GetDigitsFromString("111122")}, true},
		{"Invalid 1", args{star_04_1.GetDigitsFromString("135679")}, false},
		{"Invalid 2", args{star_04_1.GetDigitsFromString("223450")}, false},
		{"Invalid 3", args{star_04_1.GetDigitsFromString("123789")}, false},
		{"Invalid 4", args{star_04_1.GetDigitsFromString("123444")}, false},
		{"Invalid 5", args{star_04_1.GetDigitsFromString("111111")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreDigitsValid(tt.args.digits); got != tt.want {
				t.Errorf("AreDigitsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}