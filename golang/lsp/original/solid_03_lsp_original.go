// LSP - Princípio da Substituição de Liskov
// Go não tem herança, mas interfaces permitem substituição.

package main

import "fmt"

// Animal base
type Animal interface {
	Speak() string
}

// Cachorro
type Dog struct{}

func (d Dog) Speak() string {
	return "Au au!"
}

// Gato
type Cat struct{}

func (c Cat) Speak() string {
	return "Miau!"
}

func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	var a Animal

	a = Dog{}
	MakeAnimalSpeak(a)

	a = Cat{}
	MakeAnimalSpeak(a)
}
