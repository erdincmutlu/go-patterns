package main

import "fmt"

// Dependency Inversion Principle
// High Level Modules (HLM) should not depend on Low Level Modules (LLM)
// Both should depend on abstractions (interfaces in Go)

// Define Relationship between people

type Relationship int

const (
	Parent  Relationship = 1
	Child   Relationship = 2
	Sibling Relationship = 3
)

type Person struct {
	name string
	//
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level module
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent *Person, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// For Dependency Inversion
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

// high-level module
type Research struct {
	// break DIP
	// relationships Relationships

	// For Dependency Inversion
	browser RelationshipBrowser
}

// Old Way, was breaking DIP
/*
func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" &&
			rel.relationship == Parent {
			fmt.Printf("John has a child called %s\n", rel.to.name)
		}
	}
}
*/

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Printf("John has a child called %s\n", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}
