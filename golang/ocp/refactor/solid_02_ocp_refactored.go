// OCP - Refatorado em Go
// Aberto para extensão, fechado para modificação usando interfaces.
// Go favorece interfaces pequenas e funções livres.

package main

import "fmt"

// Produto
type Product struct {
	Name  string
	Price float64
}

// Interface para desconto
type DiscountStrategy interface {
	Apply(p Product) float64
}

// Desconto padrão
type DefaultDiscount struct{}

func (d DefaultDiscount) Apply(p Product) float64 {
	return p.Price * 0.9
}

// Desconto especial
type SpecialDiscount struct{}

func (s SpecialDiscount) Apply(p Product) float64 {
	return p.Price * 0.8
}

func main() {
	p := Product{Name: "Notebook", Price: 3000}
	var strategy DiscountStrategy = DefaultDiscount{}
	fmt.Printf("Preço com desconto padrão: %.2f\n", strategy.Apply(p))

	strategy = SpecialDiscount{}
	fmt.Printf("Preço com desconto especial: %.2f\n", strategy.Apply(p))
}
