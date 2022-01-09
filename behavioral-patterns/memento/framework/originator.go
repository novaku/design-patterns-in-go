package framework

type Originator struct {
	State string
}

func (e *Originator) CreateMemento() *Memento {
	return &Memento{State: e.State}
}

func (e *Originator) RestoreMemento(m *Memento) {
	e.State = m.GetSavedState()
}

func (e *Originator) SetState(state string) {
	e.State = state
}

func (e *Originator) GetState() string {
	return e.State
}