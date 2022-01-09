package framework

type Director struct {
	Builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.Builder = b
}

func (d *Director) BuildHouse() House {
	d.Builder.setDoorType()
	d.Builder.setWindowType()
	d.Builder.setNumFloor()
	return d.Builder.getHouse()
}