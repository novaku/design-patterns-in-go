package framework

import "fmt"

type Medical struct {
	Next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.Next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.Next.Execute(p)
}

func (m *Medical) SetNext(next Department) {
	m.Next = next
}