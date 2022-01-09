package main

import "design-patterns/structural-patterns/adapter/framework"

func main() {

	client := &framework.Client{}
	mac := &framework.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &framework.Windows{}
	windowsMachineAdapter := &framework.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}