// SRP - Princípio da Responsabilidade Única
// Este exemplo foi convertido de PHP para Go.
// Em Go, usamos structs e funções, evitando herança e métodos complexos.
// Comentários explicam as principais diferenças.

package main

import "fmt"

// Em Go, não usamos classes. Usamos structs para representar dados.
type User struct {
	Name  string
	Email string
}

// Função para salvar usuário (simula persistência)
func SaveUser(u User) {
	fmt.Printf("Salvando usuário: %s, email: %s\n", u.Name, u.Email)
}

// Função para enviar email (simula envio)
func SendEmail(u User, message string) {
	fmt.Printf("Enviando email para %s: %s\n", u.Email, message)
}

func main() {
	user := User{Name: "João", Email: "joao@email.com"}
	SaveUser(user)
	SendEmail(user, "Bem-vindo!")
}
