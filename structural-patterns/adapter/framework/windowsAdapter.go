package framework

import "fmt"

type WindowsAdapter struct {
	WindowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}