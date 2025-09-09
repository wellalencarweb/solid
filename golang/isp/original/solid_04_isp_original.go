// ISP - Princípio da Segregação de Interface
// Go favorece interfaces pequenas e específicas.

package main

import "fmt"

// Interface grande
type Worker interface {
	Work()
	Eat()
}

// Robô implementa só Work
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robô trabalhando")
}

// Humano implementa Work e Eat
type Human struct{}

func (h Human) Work() {
	fmt.Println("Humano trabalhando")
}

func (h Human) Eat() {
	fmt.Println("Humano comendo")
}

func main() {
	var w Worker

	w = Robot{}
	w.Work()
	// w.Eat() // erro: Robot não implementa Eat

	w = Human{}
	w.Work()
	w.Eat()
}
