package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c *Cow) Eat()   { fmt.Println("grass") }
func (c *Cow) Move()  { fmt.Println("walk") }
func (c *Cow) Speak() { fmt.Println("moo") }

func (b *Bird) Eat()   { fmt.Println("worms") }
func (b *Bird) Move()  { fmt.Println("fly") }
func (b *Bird) Speak() { fmt.Println("peep") }

func (s *Snake) Eat()   { fmt.Println("mice") }
func (s *Snake) Move()  { fmt.Println("slither") }
func (s *Snake) Speak() { fmt.Println("hsss") }

func newAnimal(kind string) (Animal, bool) {
	switch kind {
	case "cow":
		return &Cow{}, true
	case "bird":
		return &Bird{}, true
	case "snake":
		return &Snake{}, true
	default:
		return nil, false
	}
}

func main() {
	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")

		var cmd, name, arg string
		if _, err := fmt.Scan(&cmd, &name, &arg); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch cmd {
		case "newanimal":
			if _, exists := animals[name]; exists {
				fmt.Println("Animal with that name already exists.")
				continue
			}
			a, ok := newAnimal(arg)
			if !ok {
				fmt.Println("Unknown animal type:", arg)
				continue
			}
			animals[name] = a
			fmt.Println("Created it!")

		case "query":
			a, exists := animals[name]
			if !exists {
				fmt.Println("No animal named:", name)
				continue
			}
			switch arg {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Println("Unknown query:", arg)
			}

		default:
			fmt.Println("Unknown command:", cmd)
		}
	}
}