package star_08_1

import (
	"reflect"
	"testing"
)

func TestImage_DecodeImage(t *testing.T) {
	type fields struct {
		width  int
		height int
		layers []Layer
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Image
	}{
		{"Example", fields{3, 2, nil}, args{[]byte("123456789012")}, Image{3, 2, []Layer{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 0, 1, 2}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				Width:  tt.fields.width,
				Height: tt.fields.height,
				Layers: tt.fields.layers,
			}
			img.DecodeImage(tt.args.data)
			if !reflect.DeepEqual(*img, tt.want) {
				t.Errorf("(*Image)DecodeImage() = %v, want %v", img, tt.want)
			}
		})
	}
}

func TestLayer_CountInt(t *testing.T) {
	type args struct {
		digit int
	}
	tests := []struct {
		name  string
		layer Layer
		args  args
		want  int
	}{
		{"CanCount", Layer{0,1,2,0,0,2}, args{0}, 3},
		{"CanCount", Layer{0,1,2,0,0,2}, args{1}, 1},
		{"CanCount", Layer{0,1,2,0,0,2}, args{2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.layer.CountInt(tt.args.digit); got != tt.want {
				t.Errorf("CountInt() = %v, want %v", got, tt.want)
			}
		})
	}
}