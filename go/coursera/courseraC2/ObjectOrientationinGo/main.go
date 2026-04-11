package main

import (
	"fmt"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a *Animal) Eat(){
	fmt.Println(a.food)
}

func (a *Animal) Move(){
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak(){
	fmt.Println(a.noise)
}


func call(animal *Animal, method string){
	switch method {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func findAnimal(s string) *Animal {
	switch s {
	case "cow":
		return &Animal{"grass", "walk", "moo"}
	case "bird":
		return &Animal{"worms", "fly", "peep"}
	case "snake":
		return &Animal{"mice", "slither", "hsss"}
	default:
		return &Animal{}
	}

}

func main(){
	for {
		fmt.Print("> ")

		var animalName, action string
		_, err := fmt.Scanln(&animalName, &action)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		animal := findAnimal(animalName)
		call(animal, action)
	}
}