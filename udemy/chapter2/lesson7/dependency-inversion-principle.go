package main

import "fmt"

// Dependency Inversion Principle
// High-level modules should not depend on low-level modules. Both should depend on abstractions.
// Abstractions should not depend on details. Details should depend on abstractions.

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from *Person
	relationship Relationship
	to *Person
}

// low-level module
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module
// type Research struct {
// 	// break DIP
// 	relationships Relationships
// }

// func (r *Research) Investigate() {
// 	relations := r.relationships.relations
// 	for _, info := range relations {
// 		if info.from.name == "John" && info.relationship == Parent {
// 			fmt.Println("John has a child called ", info.to.name)
// 		}
// 	}
// }

// func main() {
// 	// break DIP
// 	r := Research{}
// 	// low-level part
// 	r.relationships.AddParentAndChild(&Person{"John"}, &Person{"Chris"})
// 	r.Investigate()

// }

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person{
	result := make([]*Person, 0)
	for i, info := range r.relations {
		if info.from.name == name && info.relationship == Parent {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}
// high-level module
type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called ", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)
	// r := Research{&relationships}
	// r.Investigate()
}
