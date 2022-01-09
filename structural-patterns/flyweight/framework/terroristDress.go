package framework

type TerroristDress struct {
	Color string
}

func (t *TerroristDress) GetColor() string {
	return t.Color
}

func NewTerroristDress() *TerroristDress {
	return &TerroristDress{Color: "red"}
}