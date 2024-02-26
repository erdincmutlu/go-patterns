package main

import "fmt"

type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

// Using generator
func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

// Unideomatic, kind of approach used in C++
type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("We should not be here!")
}

func main() {
	p := Person{"Alexander", "Graham", "Bell"}
	for _, name := range p.Names() {
		fmt.Println(name)
	}

	p2 := Person{"Alexander", "Graham", "Bell"}
	for name2 := range p2.NamesGenerator() {
		fmt.Println(name2)
	}

	p3 := Person{"Alexander", "", "Bell"}
	for name3 := range p3.NamesGenerator() {
		fmt.Println(name3)
	}

	p4 := Person{"Alexander", "Graham", "Bell"}
	for it := NewPersonNameIterator(&p4); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
