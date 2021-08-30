package fundamentals

import (
	"fmt"
)

type Animal interface {
	Roar() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (dog *Dog) Roar() string {
	return "Bark"
}

func (dog *Dog) NumberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (gorilla *Gorilla) Roar() string {
	return "????"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}

var _ Animal = (*Dog)(nil)     // Verify that *Dog implements Animal.
var _ Animal = (*Gorilla)(nil) // Verify that *Gorilla implements Animal.

func AnimalInfo(animal Animal) string {
	return fmt.Sprintf("This animal roar '%s' and has %d legs.", animal.Roar(), animal.NumberOfLegs())
}
