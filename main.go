package main

import (
	"fmt"
	" hery-ciaputra/demo-go-oop/contest"
	" hery-ciaputra/demo-go-oop/zoo"
)

func main() {
	animal := zoo.Animal{}
	animal.Eat()
	animal.Sleep()

	superAnimal := zoo.NewAnimal(100, 0)
	fmt.Println("SuperAnimal", superAnimal)

	cat := zoo.Cat{Animal: animal}
	cat.Eat()
	cat.Play()

	fish := zoo.NewFish{&zoo.Animal{}}
	fish.Eat()
	fish.Swim()

	fmt.Println("===============")
	contestants := []contest.Eater{
		&cat,
		&fish,
	}
	contest.DoEatingContest(contestants)
}
