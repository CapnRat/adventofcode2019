package star_08_2

import (
	"fmt"
	"github.com/CapnRat/adventofcode2019/star_08_1"
	"github.com/gookit/color"
)

const file = "star_08_1/input"

type Layer []int

type Solver struct{}

func (s *Solver) Solve() string {
	const width = 25
	const height = 6
	img := star_08_1.Image{width, height, nil}
	img.DecodeImage(star_08_1.ReadInput(file))

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			for _, layer := range img.Layers {
				layerIndex := y * img.Width + x
				pixel := layer[layerIndex]
				if pixel == 0 {
					color.BgRed.Print(" ")
					break
				}
				if pixel == 1 {
					color.Green.Print("█")
					break
				}
			}
		}
		fmt.Print("\r\n")
	}

	return "Done"
}