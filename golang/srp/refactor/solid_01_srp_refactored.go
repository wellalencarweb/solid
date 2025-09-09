// SRP - Refatorado em Go
// Separação clara das responsabilidades usando structs e funções.
// Go favorece composição e funções livres ao invés de métodos em classes.

package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type UserRepository struct{}

func (r UserRepository) Save(u User) {
	fmt.Printf("Salvando usuário: %s, email: %s\n", u.Name, u.Email)
}

type EmailService struct{}

func (e EmailService) Send(u User, message string) {
	fmt.Printf("Enviando email para %s: %s\n", u.Email, message)
}

func main() {
	user := User{Name: "João", Email: "joao@email.com"}
	repo := UserRepository{}
	email := EmailService{}

	repo.Save(user)
	email.Send(user, "Bem-vindo!")
}
