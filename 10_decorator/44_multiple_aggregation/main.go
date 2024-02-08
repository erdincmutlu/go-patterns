package main

import "fmt"

// Initial Design
// Dragon from bird and lizard structures
type BirdOld struct {
	Age int
}

func (b *BirdOld) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying!")
	}
}

type LizardOld struct {
	Age int
}

func (l *LizardOld) Crawl() {
	if l.Age < 10 {
		fmt.Println("Crawling!")
	}
}

// Both Bird and Lizard has Age fields. This will cause problem
type DragonOld struct {
	BirdOld
	LizardOld
}

func (d *DragonOld) AgeOld() int {
	return d.BirdOld.Age
}

func (d *DragonOld) SetAgeOld(age int) {
	d.BirdOld.Age = age
	d.LizardOld.Age = age
}

// Using Decorator
type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Printf("Flying!")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int       { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling!")
	}
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) Age() int {
	return d.bird.age
}

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}
func main() {
	// Initial Design
	dOld := DragonOld{}
	// d.Age = 10 // Not allowed. Ambiguous selector d.Age

	// dOld.BirdOld.Age = 10
	// dOld.LizardOld.Age = 10
	dOld.SetAgeOld(10)
	dOld.BirdOld.Age = 55 // You are breking the consistency
	dOld.Fly()
	dOld.Crawl()

	// Using Decorator
	d := Dragon{}
	d.SetAge(10)
	d.Fly()
	d.Crawl()
}
