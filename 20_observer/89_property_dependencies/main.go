package main

import (
	"container/list"
	"fmt"
)

// Observable, Observer
// Patient, doctor

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type PropertyChange struct {
	Name  string // "Age" "Height"
	Value interface{}
}

type Person struct {
	Observable
	age int
} // Age() SetAge()

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return

	}

	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct{}

func (e *ElectoralRoll) Notify(data interface{}) {
	pc, ok := data.(PropertyChange)
	if ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func main() {
	p := NewPerson(0)
	er := &ElectoralRoll{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Printf("Setting age to %d\n", i)
		p.SetAge(i)
	}
}
