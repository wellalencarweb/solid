// DIP - Refatorado em Go
// Serviço depende de abstração (interface), não de implementação concreta.

package main

import "fmt"

// Interface de repositório
type Repository interface {
	Save(data string)
}

// Implementação MySQL
type MySQLRepository struct{}

func (r MySQLRepository) Save(data string) {
	fmt.Printf("Salvando '%s' no MySQL\n", data)
}

// Serviço depende da interface
type UserService struct {
	Repo Repository
}

func (s UserService) Register(name string) {
	s.Repo.Save(name)
}

func main() {
	repo := MySQLRepository{}
	service := UserService{Repo: repo}
	service.Register("Maria")
}
