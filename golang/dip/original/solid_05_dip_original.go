// DIP - Princípio da Inversão de Dependência
// Go usa interfaces para desacoplar dependências.

package main

import "fmt"

// Repositório concreto
type MySQLRepository struct{}

func (r MySQLRepository) Save(data string) {
	fmt.Printf("Salvando '%s' no MySQL\n", data)
}

// Serviço depende do repositório concreto
type UserService struct {
	Repo MySQLRepository
}

func (s UserService) Register(name string) {
	s.Repo.Save(name)
}

func main() {
	repo := MySQLRepository{}
	service := UserService{Repo: repo}
	service.Register("Maria")
}
