package star_08_1

import (
	"fmt"
	"github.com/psampaz/slice"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

const file = "star_08_1/input"

type Layer []int

type Image struct{
	width int
	height int
	layers []Layer
}

type Solver struct{}

func (s *Solver) Solve() string {
	const width = 25
	const height = 6
	img := Image{width, height, nil}
	img.DecodeImage(ReadInput(file))

	type zeros struct {
		count int
		layer int
	}
	z := zeros{math.MaxInt64, -1}
	for l := 0; l < len(img.layers); l++{
		zeroCount := img.layers[l].CountInt(0)
		if zeroCount < z.count {
			z.count = zeroCount
			z.layer = l
		}
	}

	ones := img.layers[z.layer].CountInt(1)
	twos := img.layers[z.layer].CountInt(2)

	return strconv.Itoa(ones * twos)
}

func (layer Layer) CountInt(digit int) int {
	layerCopy := append(Layer(nil), layer...)
	return len(slice.FilterInt(layerCopy, func(x int) bool { return x == digit }))
}

func (img *Image) DecodeImage(data []byte) {
	if len(data) % (img.width * img.height) != 0 {
		panic("data does not fit into image")
	}

	numLayers := len(data) / (img.width * img.height)

	img.layers = make([]Layer, numLayers)
	dataIndex := 0
	for l := 0; l < numLayers; l++ {
		layer := make(Layer, img.width * img.height)
		for y := 0; y < img.height; y++ {
			for x := 0; x < img.width; x++ {
				layerIndex := y * img.width + x
				layer[layerIndex], _ = strconv.Atoi(string(data[dataIndex]))
				dataIndex++
			}
		}
		img.layers[l] = layer
	}
}

func ReadInput(file string) []byte {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	return bytes
}