package framework


type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shirt struct {
	Logo string
	Size int
}

func (s *Shirt) SetLogo(logo string) {
	s.Logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.Logo
}

func (s *Shirt) SetSize(size int) {
	s.Size = size
}

func (s *Shirt) GetSize() int {
	return s.Size
}