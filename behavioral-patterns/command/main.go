package main

import "design-patterns/behavioral-patterns/command/framework"

func main() {
	tv := &framework.Tv{}

	onCommand := &framework.OnCommand{
		Device: tv,
	}

	offCommand := &framework.OffCommand{
		Device: tv,
	}

	onButton := &framework.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &framework.Button{
		Command: offCommand,
	}
	offButton.Press()
}