// OCP - Princípio do Aberto/Fechado
// Exemplo convertido de PHP para Go.
// Go não tem herança, então usamos interfaces e composição.

package main

import "fmt"

// Produto simples
type Product struct {
	Name  string
	Price float64
}

// Calcula preço com desconto (sem extensão)
func CalculateDiscount(p Product) float64 {
	return p.Price * 0.9 // 10% de desconto fixo
}

func main() {
	p := Product{Name: "Notebook", Price: 3000}
	discount := CalculateDiscount(p)
	fmt.Printf("Preço com desconto: %.2f\n", discount)
}
