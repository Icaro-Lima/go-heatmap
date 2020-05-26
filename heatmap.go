package heatmap

import "github.com/Icaro-Lima/go-heatmap/stamp"

type Heatmap struct {
	Buffer        []float32
	Max           float32
	Width, Height int
}

var defaultStamp = stamp.New(9)

func New(width, height int) *Heatmap {
	h := Heatmap{
		Width:  width,
		Height: height,
		Buffer: make([]float32, width*height),
	}

	return &h
}

func (h *Heatmap) AddPoint(x int, y int) {
	h.AddPointWithStamp(x, y, defaultStamp)
}

func (h *Heatmap) AddPointWithStamp(x int, y int, stamp *stamp.Stamp) {
	if x >= h.Width || y >= h.Height {
		return
	}

	var x0 int
	var y0 int
	var x1 int
	var y1 int
	if x0 = 0; x < stamp.Width/2 {
		x0 = stamp.Width/2 - x
	}
	if y0 = 0; y < stamp.Height/2 {
		y0 = stamp.Height/2 - y
	}
	if x1 = stamp.Width/2 + (h.Width - x); (x + stamp.Width/2) < h.Width {
		x1 = stamp.Width
	}
	if y1 = stamp.Height/2 + (h.Height - y); (y + stamp.Height/2) < h.Height {
		y1 = stamp.Height
	}

	for iy := y0; iy < y1; iy++ {
		line := ((y+iy)-stamp.Height/2)*h.Width + (x + x0) - stamp.Width/2
		stampLine := iy*stamp.Width + x0

		for ix := x0; ix < x1; ix, line, stampLine = ix+1, line+1, stampLine+1 {
			line += stampLine
			if h.Buffer[line] > h.Max {
				h.Max = h.Buffer[line]
			}
		}
	}
}
