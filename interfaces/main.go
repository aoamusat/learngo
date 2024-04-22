package main

import "fmt"

// We start by defining our Animal interface:
type Animal interface {
	Speak() string
}

type Bird interface {
	Fly(destination string) string
}

type Eagle struct {
	name string
}

func (bird Eagle) Fly(destination string) string {
	return fmt.Sprintf("%s flying to %s", bird.name, destination)
}

func (bird Eagle) Speak() string {
	return fmt.Sprintf("Eagle %s Speaking... Kro Kro", bird.name)
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}, Eagle{"Shaksha"}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	example()
}
