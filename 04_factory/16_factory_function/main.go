package main

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

// Factory Function: Freestanding function which return an instance of the
// struct that you want to create
func NewPerson(name string, age int) *Person {
	// Have some additional logic
	if age < 16 {
		// ...
	}

	return &Person{name, age, 2}
}

func main() {
	// Direct
	// p := Person{"John", 22, 1}

	// Using factory
	p := NewPerson("John", 33)
	p.EyeCount = 1
}
