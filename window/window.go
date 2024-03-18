package window

type window struct {
	Width  int
	Height int
}

func NewWindow(width, height int) *window {
	return &window{width, height}
}