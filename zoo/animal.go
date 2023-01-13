package zoo

import "fmt"

type Animal struct {
	energy, hungry int
}

func NewAnimal(energy, hungry int) *Animal {
	return &Animal{
		energy: energy,
		hungry: hungry,
	}
}

// Eat method function
func (a *Animal) Eat() {
	a.hungry++
	a.energy++
	fmt.Println("Animal eating...")
}

func (a *Animal) Sleep() {
	a.hungry++
	a.energy++
	fmt.Println("Animal sleeping...")
}
