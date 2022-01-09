package main

import (
	"fmt"

	"design-patterns/structural-patterns/bridge/framework"
)

func main() {

	hpPrinter := &framework.Hp{}
	epsonPrinter := &framework.Epson{}

	macComputer := &framework.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &framework.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}