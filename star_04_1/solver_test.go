package star_04_1

import (
	"reflect"
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
		{"Valid 1", args{GetDigitsFromString("122345")}, true},
		{"Valid 2", args{GetDigitsFromString("111123")}, true},
		{"Valid 3", args{GetDigitsFromString("111111")}, true},
		{"Invalid 1", args{GetDigitsFromString("135679")}, false},
		{"Invalid 2", args{GetDigitsFromString("223450")}, false},
		{"Invalid 3", args{GetDigitsFromString("123789")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreDigitsValid(tt.args.digits); got != tt.want {
				t.Errorf("AreDigitsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDigitsFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{"111111"}, []int{1,1,1,1,1,1}},
		{"Example 2", args{"0123456"}, []int{0,1,2,3,4,5,6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDigitsFromString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDigitsFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}