package main

import "fmt"

//Creature type
type Creature struct {
	Species string
}

// Reset type
func (c *Creature) Reset() {
	c.Species = ""
}

func main() {
	var creature *Creature = &Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature)
	// changeCreature(creature) //Pass by value
	// changeCreature(creature) //Pass by reference
	creature.Reset()
	fmt.Printf("2) %+v\n", creature)
}

// func changeCreature(creature Creature) { //Pass by value
func changeCreature(creature *Creature) { //Pass by reference
	if creature == nil {
		fmt.Println("creature is nil")
		return
	}
	creature.Species = "Jellyfish"
	fmt.Printf("2) %+v\n", creature)
}
