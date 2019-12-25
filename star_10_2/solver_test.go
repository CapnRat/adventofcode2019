package star_10_2

import (
	"math"
	"reflect"
	"testing"
)

func TestEncodeSpacePoint(t *testing.T) {
	type args struct {
		point SpacePoint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{SpacePoint{8, 2}}, 802},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeSpacePoint(tt.args.point); got != tt.want {
				t.Errorf("EncodeSpacePoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveWithInputFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want SpacePoint
	}{
		{"Example4", args{"../star_10_1/example4"}, SpacePoint{8,2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveWithInputFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveWithInputFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortField(t *testing.T) {
	type args struct {
		field         AsteroidField
		deathAsteroid Asteroid
	}
	tests := []struct {
		name string
		args args
		want AsteroidField
	}{
		{
			"DeathAsteroidIsExcluded",
			args {
				AsteroidField {
					SpacePoint{0,0},
					SpacePoint{1,1},
				},
				Asteroid {
					Point:            SpacePoint{0,0},
					VisibleAsteroids: 1,
				},
			},
			AsteroidField {
				SpacePoint{1,1},
			},
		},
		{
			"SameRadialIsSorted",
			args {
				AsteroidField {
					SpacePoint{0,0},
					SpacePoint{0,2},
					SpacePoint{0,1},
				},
				Asteroid {
					Point:            SpacePoint{0,0},
					VisibleAsteroids: 1,
				},
			},
			AsteroidField {
				SpacePoint{0,1},
				SpacePoint{0,2},
			},
		},
		{
			"RadialSortedBeforeSame",
			args {
				AsteroidField {
					SpacePoint{0,0},
					SpacePoint{1,1},
					SpacePoint{0,2},
					SpacePoint{0,1},
				},
				Asteroid {
					Point:            SpacePoint{0,0},
					VisibleAsteroids: 1,
				},
			},
			AsteroidField {
				SpacePoint{1,1},
				SpacePoint{0,1},
				SpacePoint{0,2},
			},
		},
		{
			"RadialsAreSorted",
			args {
				AsteroidField {
					SpacePoint{ 1, 0},
					SpacePoint{-1, 0},
					SpacePoint{ 0, 0},
					SpacePoint{ 0, 1},
					SpacePoint{ 0,-1},
					SpacePoint{ 1, 1},
					SpacePoint{ 1,-1},
					SpacePoint{-1,-1},
					SpacePoint{-1, 1},
				},
				Asteroid {
					Point:            SpacePoint{0,0},
					VisibleAsteroids: 1,
				},
			},
			AsteroidField {
				SpacePoint{ 0,-1},
				SpacePoint{ 1,-1},
				SpacePoint{ 1, 0},
				SpacePoint{ 1, 1},
				SpacePoint{ 0, 1},
				SpacePoint{-1, 1},
				SpacePoint{-1, 0},
				SpacePoint{-1,-1},
			},
		},
		{
			"RadialsStaySorted",
			args {
				AsteroidField {
					SpacePoint{ 0, 1},
					SpacePoint{ 1, 1},
					SpacePoint{ 1, 0},
					SpacePoint{ 1,-1},
					SpacePoint{ 0,-1},
					SpacePoint{-1,-1},
					SpacePoint{-1, 0},
					SpacePoint{-1, 1},
					SpacePoint{ 0, 0},
				},
				Asteroid {
					Point:            SpacePoint{0,0},
					VisibleAsteroids: 1,
				},
			},
			AsteroidField {
				SpacePoint{ 0,-1},
				SpacePoint{ 1,-1},
				SpacePoint{ 1, 0},
				SpacePoint{ 1, 1},
				SpacePoint{ 0, 1},
				SpacePoint{-1, 1},
				SpacePoint{-1, 0},
				SpacePoint{-1,-1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortField(tt.args.field, tt.args.deathAsteroid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngleBetweenRadials(t *testing.T) {
	type args struct {
		i Radial
		j Radial
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"ParallelRadials",
			args {
				Radial{0, -1},
				Radial{0, -1},
			},
			0,
		},
		{
			"RightPerpendicularRadials",
			args {
				Radial{0, -1},
				Radial{1, 0},
			},
			math.Pi / 2,
		},
		{
			"LeftPerpendicularRadials",
			args {
				Radial{ 0, -1},
				Radial{-1, 0},
			},
			3 * math.Pi / 2,
		},
		{
			"OppositeRadials",
			args {
				Radial{0,-1},
				Radial{0, 1},
			},
			math.Pi,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AngleBetweenRadials(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("AngleBetweenRadials() = %v, want %v", got, tt.want)
			}
		})
	}
}