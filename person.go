package main

import "fmt"

// type deck []string
/*
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, "=>", card)
	}
}*/

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo
}

func (p *person) updateName(firstName string) {
	// p.firstName = firstName // Works
	(*p).firstName = firstName // Works
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (c contactInfo) print() {
	fmt.Println(c)
}
