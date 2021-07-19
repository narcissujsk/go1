package test1

type Circle struct {
	Radius float64
}

func (c Circle) GetArea() float64 {
	return 3.14 * c.Radius * c.Radius
	return 0
}
