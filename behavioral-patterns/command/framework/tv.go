package framework

import "fmt"

type Tv struct {
	IsRunning bool
}

func (t *Tv) On() {
	t.IsRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) Off() {
	t.IsRunning = false
	fmt.Println("Turning tv off")
}