package star_01_1

import "testing"

func TestSolver_CalculateFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"12", args{12}, 2},
		{"14", args{14}, 2},
		{"1969", args{1969}, 654},
		{"100756", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{}
			if got := s.CalculateFuel(tt.args.mass); got != tt.want {
				t.Errorf("CalculateFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
