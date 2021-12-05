package main

// OCP
// open for extension, closed for modification
// Specification is an interface

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)	

type Product struct {
	name string
	color Color
	size Size
}

type Filter struct {
}

func (f *Filter) FilterByColor (products []Product, color Color) []*Product {
	var result = make([]*Product, 0)
	for i, p := range products {
		if p.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}
 
func (f *Filter) FilterBySize (products []Product, size Size) []*Product {
	var result = make([]*Product, 0)
	for i, p := range products {
		if p.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor (products []Product, size Size, color Color) []*Product {
	var result = make([]*Product, 0)
	for i, p := range products {
		if p.size == size && p.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// specification pattern
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func(spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}
type SizeSpecification struct {
	size Size
}

func(s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct {
}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	var result = make([]*Product, 0)
	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}
	return result
}

// composite specification pattern
type AndSpecification struct {
	first, second Specification
}

func(a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

// func main() {
// 	apple := Product{"Apple", green, small}
// 	tree := Product{"Tree", green, large}
// 	house := Product{"House", blue, large}
// 	products := []Product{apple, tree, house}
// 	fmt.Println("Green products:")
// 	f := Filter{}
// 	for _, p := range f.FilterByColor(products, green) {
// 		fmt.Printf(" - %s is green\n", p.name)
// 	}
// }

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}
	products := []Product{apple, tree, house}
	fmt.Println("Green products:")
	f := BetterFilter{}	
	greepSpec := ColorSpecification{green}
	for _, p := range f.Filter(products, greepSpec) {
		fmt.Printf(" - %s is green\n", p.name)
	}
	largeSpec := SizeSpecification{large}
	largeGreenSpec := AndSpecification{greepSpec, largeSpec}
	for _, p := range f.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", p.name)
	}
}