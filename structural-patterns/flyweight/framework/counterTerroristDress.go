package framework

type CounterTerroristDress struct {
	Color string
}

func (c *CounterTerroristDress) GetColor() string {
	return c.Color
}

func NewCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{Color: "green"}
}