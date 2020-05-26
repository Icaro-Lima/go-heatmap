package stamp

import "math"

type Stamp struct {
	Buffer        []float32
	Width, Height int
}

type DistributionShape func(distribution float32) float32

func New(radius int) *Stamp {
	return NewNonlinear(radius, linearDistributionShape)
}

func NewNonlinear(radius int, distributionShape DistributionShape) *Stamp {
	d := 2*radius + 1

	stamp := Stamp{
		Buffer: make([]float32, d*d),
		Width:  d,
		Height: d,
	}

	for y := 0; y < d; y++ {
		line := y * d
		for x := 0; x < d; x, line = x+1, line+1 {
			dist := float32(math.Sqrt(float64((x-radius)*(x-radius)+(y-radius)*(y-radius)))) / float32(radius+1)
			ds := distributionShape(dist)
			var clampedDs float32
			if ds > 1 {
				clampedDs = 1
			} else if ds < 0 {
				clampedDs = 0
			} else {
				clampedDs = ds
			}

			stamp.Buffer[line] = 1 - clampedDs
		}
	}

	return &stamp
}

func linearDistributionShape(dist float32) float32 {
	return dist
}
