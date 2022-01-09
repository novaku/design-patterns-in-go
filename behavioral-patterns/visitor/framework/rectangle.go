package framework

type Rectangle struct {
	L int
	B int
}

func (t *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(t)
}

func (t *Rectangle) GetType() string {
	return "rectangle"
}