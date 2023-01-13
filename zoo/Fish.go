package zoo

import "fmt"

type Fish struct {
	animal *Animal
}

func NewFish(animal *Animal) *Fish {
	return &Fish{animal}
}

func (f *Fish) Eat() {
	f.animal.Eat()
	f.animal.energy++
	fmt.Println("Fish eating...")
}

func (f *Fish) Swim() {
	f.animal.energy--
	f.animal.hungry++
	fmt.Println("Fish swimming...")
}
