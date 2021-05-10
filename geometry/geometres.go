package geometry

type Triangle struct {
	size int
}

func (t *Triangle) DoubleSize() {
	t.size *= 2
}
func (t *Triangle) SetSize(size int) {
	t.size = size
}
func (t *Triangle) Perimeter() int {
	t.DoubleSize()
	return t.size * 3
}
func (t *Triangle) name() int {
	return t.size * 2
}
