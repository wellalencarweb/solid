// LSP - Refatorado em Go
// Interfaces garantem substituição sem herança.
// Go não tem métodos virtuais, mas permite polimorfismo por interface.

package main

import "fmt"

// Animal
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

// Pato (nova implementação)
type Duck struct{}

func (d Duck) Speak() string {
	return "Quack!"
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

	a = Duck{}
	MakeAnimalSpeak(a)
}
