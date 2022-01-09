package framework

import "fmt"

type Reception struct {
	Next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.Next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.Next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.Next = next
}