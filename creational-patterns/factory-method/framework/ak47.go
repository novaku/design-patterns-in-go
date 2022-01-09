package framework

type Ak47 struct {
	Gun
}

func (a Ak47) SetName(name string) {
	a.Name = name
}

func (a Ak47) SetPower(power int) {
	a.Power = power
}

func (a Ak47) GetName() string {
	return a.Name
}

func (a Ak47) GetPower() int {
	return a.Power
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			Name:  "AK47 gun",
			Power: 4,
		},
	}
}