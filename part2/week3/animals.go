package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct {
	Food string
	Locomotion string
	Noise string
}

func (a *Animal) Eat() string {
	return a.Food
}

func (a *Animal) Move() string {
	return a.Locomotion
}

func (a *Animal) Speak() string {
	return a.Noise
}

func RaiseError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetAnimal(name string) Animal {
	switch name {
	case "cow": return Animal{Food: "grass", Locomotion: "walk", Noise: "moo"}
	case "bird": return Animal{Food: "worms", Locomotion: "fly", Noise: "peep"}
	case "snake": return Animal{Food: "mice", Locomotion: "slither", Noise: "hsss"}
	default: return Animal{}
	}
}

func main() {
	fmt.Println("Press Control-C to exit.")
	fmt.Println("Enter an animal and associated verb. Example: cow eat")

	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("> ")
		input, err := reader.ReadString('\n')
		RaiseError(err)
		
		input = strings.Trim(input, "\n")
		inputs := strings.Fields(input)
		animal := GetAnimal(inputs[0])

		if len(inputs) == 1 {
			fmt.Println("An animal and a verb must be specified.")
			continue
		}

		if len(animal.Eat()) == 0 {
			fmt.Println("Invalid animal. Value:", inputs[0], ": Try: cow, bird, or snake.")
			continue
		}

		switch inputs[1] {
		case "eat": fmt.Println(animal.Eat())
		case "move": fmt.Println(animal.Move())
		case "speak": fmt.Println(animal.Speak())
		default: fmt.Println("Invalid verb for", inputs[0], ": Try: eat, move, speak")
		}
	}
}