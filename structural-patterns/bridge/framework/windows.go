package framework

import "fmt"

type Windows struct {
	printer printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p printer) {
	w.printer = p
}