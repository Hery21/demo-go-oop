package zoo

import "fmt"

type Cat struct {
	Animal //embedding
}

func (c *Cat) meow() {
	fmt.Println("Meow...")
}

func (c *Cat) Eat() {
	c.Animal.Eat()
	c.meow()
}

func (c *Cat) Play() {
	c.energy++
	c.hungry++
	c.meow()
	fmt.Println("Cat Playing...")
}
