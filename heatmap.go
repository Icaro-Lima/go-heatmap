package heatmap

type Heatmap struct {
	Buffer []float32
	Max    float32
	Width  int
	Height int
}

func (h *Heatmap) New(width, height int) {
	h.Width = width
	h.Height = height
	h.Buffer = make([]float32, width*height)
}
