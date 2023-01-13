package testing

type Rectangle struct {
	width, length int
}

func (rect *Rectangle) Area() int {
	return rect.width * rect.length
}

func (rect *Rectangle) Perimeter() int {
	return 2 * (rect.length + rect.width)
}
