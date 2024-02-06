package main

import "fmt"

type Address struct {
	StreetAddress string
	City          string
	Country       string
}
type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John", &Address{"123 London Road", "London", "UK"}}

	// This is wrong
	// jane := john
	// jane.Name = "Jane" // Ok
	// jane.Address.StreetAddress = "321 Baker St"

	// deep copying
	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
