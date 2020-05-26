package heatmap

import (
	"github.com/Icaro-Lima/go-heatmap/stamp"
	"image"
	"image/color"
	"image/color/palette"
)

type Heatmap struct {
	Buffer        []float32
	Max           float32
	Width, Height int
}

var defaultStamp = stamp.New(4)

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
			h.Buffer[line] += stamp.Buffer[stampLine]
			if h.Buffer[line] > h.Max {
				h.Max = h.Buffer[line]
			}
		}
	}
}

func (h *Heatmap) RenderDefault() *image.NRGBA {
	return h.Render(palette.Plan9)
}

func (h *Heatmap) Render(palette color.Palette) *image.NRGBA {
	var saturation float32
	if saturation = 1; h.Max > 0 {
		saturation = h.Max
	}

	return h.RenderSaturated(palette, saturation)
}

func (h *Heatmap) RenderSaturated(palette color.Palette, saturation float32) *image.NRGBA {
	output := image.NewNRGBA(image.Rect(0, 0, h.Width, h.Height))

	for y := 0; y < h.Height; y++ {
		buffLine := y * h.Width

		for x := 0; x < h.Width; x, buffLine = x+1, buffLine+1 {
			var val float32
			if val = h.Buffer[buffLine]; h.Buffer[buffLine] > saturation {
				val = saturation
			}
			val /= saturation

			idx := (int)((float32)(len(palette)-1)*val + 0.5)

			output.Set(x, y, palette[idx])
		}
	}

	return output
}
