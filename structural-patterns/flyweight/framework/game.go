package framework

type Game struct {
	Terrorists        []*Player
	CounterTerrorists []*Player
}

func NewGame() *Game {
	return &Game{
		Terrorists:        make([]*Player, 1),
		CounterTerrorists: make([]*Player, 1),
	}
}

func (c *Game) AddTerrorist(dressType string) {
	player := NewPlayer("T", dressType)
	c.Terrorists = append(c.Terrorists, player)
	return
}

func (c *Game) AddCounterTerrorist(dressType string) {
	player := NewPlayer("CT", dressType)
	c.CounterTerrorists = append(c.CounterTerrorists, player)
	return
}