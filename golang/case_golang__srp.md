# 🧪 Caso Prático: Single Responsibility Principle em Go

## Contexto

No Go, o SRP é aplicado separando responsabilidades em structs e funções distintas, evitando métodos que misturam lógica de negócio, persistência e comunicação.

## Exemplo Problemático

```go
type Order struct {
	items []string
	prices []float64
}

func (o *Order) AddItem(item string, price float64) {
	o.items = append(o.items, item)
	o.prices = append(o.prices, price)
}

func (o *Order) CalculateTotal() float64 {
	total := 0.0
	for _, price := range o.prices {
		total += price
	}
	return total
}

// 🚨 Mistura responsabilidades!
func (o *Order) ProcessPayment() {}
func (o *Order) SendEmail() {}
func (o *Order) SaveToDatabase() {}
```

## Exemplo Corrigido

```go
type OrderSRP struct {
	items []string
	prices []float64
}

func (o *OrderSRP) AddItem(item string, price float64) {
	o.items = append(o.items, item)
	o.prices = append(o.prices, price)
}

func (o *OrderSRP) CalculateTotal() float64 {
	total := 0.0
	for _, price := range o.prices {
		total += price
	}
	return total
}

// Responsabilidade separada: pagamento
 type PaymentProcessor struct{}
func (p PaymentProcessor) ProcessPayment(o *OrderSRP) {}

// Responsabilidade separada: email
 type EmailService struct{}
func (e EmailService) SendOrderConfirmation(o *OrderSRP) {}

// Responsabilidade separada: persistência
 type OrderRepository struct{}
func (r OrderRepository) Save(o *OrderSRP) {}
```

## Benefícios
- Código mais coeso e testável
- Facilidade de manutenção
- Menos conflitos em equipe

## Execução

```go
func main() {
	order := &OrderSRP{}
	order.AddItem("Produto X", 200)
	order.AddItem("Produto Y", 80)

	payment := PaymentProcessor{}
	email := EmailService{}
	repo := OrderRepository{}

	payment.ProcessPayment(order)
	email.SendOrderConfirmation(order)
	repo.Save(order)
}
```
