package main

import "design-patterns/behavioral-patterns/chain-of-responsibility/framework"

func main() {

	cashier := &framework.Cashier{}

	//Set next for medical department
	medical := &framework.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &framework.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &framework.Reception{}
	reception.SetNext(doctor)

	patient := &framework.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}