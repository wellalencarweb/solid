// ISP - Refatorado em Go
// Interfaces segregadas para cada responsabilidade.

package main

import "fmt"

// Interface de trabalho
type Workable interface {
	Work()
}

// Interface de alimentação
type Eatable interface {
	Eat()
}

// Robô só trabalha
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robô trabalhando")
}

// Humano trabalha e come
type Human struct{}

func (h Human) Work() {
	fmt.Println("Humano trabalhando")
}

func (h Human) Eat() {
	fmt.Println("Humano comendo")
}

func main() {
	var w Workable = Robot{}
	w.Work()

	var h Human = Human{}
	h.Work()
	h.Eat()
}
