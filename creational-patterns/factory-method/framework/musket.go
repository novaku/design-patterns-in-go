package framework

type Musket struct {
	Gun
}

func (m Musket) SetName(name string) {
	m.Name = name
}

func (m Musket) SetPower(power int) {
	m.Power = power
}

func (m Musket) GetName() string {
	return m.Name
}

func (m Musket) GetPower() int {
	return m.Power
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			Name:  "Musket gun",
			Power: 1,
		},
	}
}