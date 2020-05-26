package heatmap

type Heatmap struct {
	Buffer        []float32
	Max           float32
	Width, Height int
}

func New(width, height int) *Heatmap {
	h := Heatmap{
		Width:  width,
		Height: height,
		Buffer: make([]float32, width*height),
	}

	return &h
}
