
# 🎯 SOLID com PHP

---

## 🎬 Introdução ao SOLID

**📋 O que é SOLID?**
- Princípios de design orientado a objetos
- Criados por Robert C. Martin (Uncle Bob)
- Acrônimo para cinco princípios fundamentais

**🎯 Por que usar SOLID?**
- Código mais limpo e manutenível
- Redução de acoplamento
- Facilita testes e extensibilidade
- Menos "code smells" e dívida técnica

**📌 Agenda:**
1. S - Single Responsibility Principle
2. O - Open/Closed Principle  
3. L - Liskov Substitution Principle
4. I - Interface Segregation Principle
5. D - Dependency Inversion Principle

---

## ⚡ Single Responsibility Principle

**📖 Conceito:**
"Uma classe deve ter apenas uma razão para mudar"

**❌ Exemplo Problemático:**
```php
class Order {
    public function calculateTotal() {}
    public function processPayment() {}
    public function sendEmail() {}
    public function saveToDatabase() {}
}
```

**✅ Exemplo Corrigido:**
```php
class Order {
    public function calculateTotal() {}
}

class PaymentProcessor {
    public function processPayment(Order $order) {}
}

class EmailService {
    public function sendOrderConfirmation(Order $order) {}
}

class OrderRepository {
    public function save(Order $order) {}
}
```

**🎁 Benefícios:**
- Classes mais coesas e focadas
- Facilidade de teste unitário
- Menos conflitos em trabalho em equipe
- Manutenção simplificada

---

## 🔓 Open/Closed Principle

**📖 Conceito:**
"Entidades devem estar abertas para extensão, mas fechadas para modificação"

**❌ Exemplo Problemático:**
```php
class PaymentProcessor {
    public function process($paymentType) {
        if ($paymentType === 'credit_card') {
            // process credit card
        } elseif ($paymentType === 'paypal') {
            // process paypal
        }
    }
}
```

**✅ Exemplo Corrigido:**
```php
interface PaymentMethod {
    public function process();
}

class CreditCardPayment implements PaymentMethod {
    public function process() {}
}

class PayPalPayment implements PaymentMethod {
    public function process() {}
}

class PaymentProcessor {
    public function processPayment(PaymentMethod $paymentMethod) {
        $paymentMethod->process();
    }
}
```

**🎁 Benefícios:**
- Fácil adição de novas funcionalidades
- Menos regressões em código existente
- Arquitetura extensível e modular

---

## 🔄 Liskov Substitution Principle

**📖 Conceito:**
"Subtipo deve ser substituível por seu tipo base sem alterar o comportamento"

**❌ Exemplo Problemático:**
```php
class Bird {
    public function fly() {}
}

class Penguin extends Bird {
    public function fly() {
        throw new Exception("Penguins can't fly!");
    }
}
```

**✅ Exemplo Corrigido:**
```php
interface FlyingBird {
    public function fly();
}

interface SwimmingBird {
    public function swim();
}

class Eagle implements FlyingBird {
    public function fly() {}
}

class Penguin implements SwimmingBird {
    public function swim() {}
}
```

**🎁 Benefícios:**
- Polimorfismo confiável e previsível
- Prevenção de surpresas desagradáveis
- Herança semanticamente correta

---

## 🧩 Interface Segregation Principle

**📖 Conceito:**
"Muitas interfaces específicas são melhores que uma interface geral"

**❌ Exemplo Problemático:**
```php
interface Worker {
    public function work();
    public function eat();
    public function sleep();
}

class Robot implements Worker {
    public function work() {}
    public function eat() { /* Robots don't eat! */ }
    public function sleep() { /* Robots don't sleep! */ }
}
```

**✅ Exemplo Corrigido:**
```php
interface Workable {
    public function work();
}

interface Eatable {
    public function eat();
}

interface Sleepable {
    public function sleep();
}

class Human implements Workable, Eatable, Sleepable {
    public function work() {}
    public function eat() {}
    public function sleep() {}
}

class Robot implements Workable {
    public function work() {}
}
```

**🎁 Benefícios:**
- Interfaces enxutas e específicas
- Sem métodos forçados ou vazios
- Melhor organização e coesão

---

## 🔌 Slide 6: Dependency Inversion Principle

**📖 Conceito:**
"Dependa de abstrações, não de implementações concretas"

**❌ Exemplo Problemático:**
```php
class MySQLConnection {
    public function connect() {}
}

class UserRepository {
    private $connection;
    
    public function __construct() {
        $this->connection = new MySQLConnection();
    }
}
```

**✅ Exemplo Corrigido:**
```php
interface DatabaseConnection {
    public function connect();
}

class MySQLConnection implements DatabaseConnection {
    public function connect() {}
}

class PostgreSQLConnection implements DatabaseConnection {
    public function connect() {}
}

class UserRepository {
    private $connection;
    
    public function __construct(DatabaseConnection $connection) {
        $this->connection = $connection;
    }
}
```

**🎁 Benefícios:**
- Baixo acoplamento entre componentes
- Fácil substituição de implementações
- Melhor testabilidade com mocks
- Arquitetura flexível e sustentável

---

## 🧪 Caso Prático

**✅ Exemplo:**
```php
class OrderService {
    public function __construct(
        private PaymentMethod $paymentMethod,
        private NotificationService $notification
    ) {}
    
    public function processOrder(Order $order) {
        $this->paymentMethod->process($order);
        $this->notification->send($order);
    }
}
```

**🎁 Benefícios Combinados:**
- Código altamente testável
- Fácil manutenção e evolução
- Baixo acoplamento entre componentes
- Escalabilidade e flexibilidade

---

## 🏁 Conclusão

**📚 Resumo dos Princípios:**
- S - Single Responsibility: Uma classe, uma responsabilidade
- O - Open/Closed: Estenda, não modifique
- L - Liskov Substitution: Substituição segura
- I - Interface Segregation: Interfaces específicas
- D - Dependency Inversion: Dependa de abstrações

**💡 Recomendações Práticas:**
- Aplique gradualmente em projetos existentes
- Evite over-engineering
- Foque em pontos de maior complexidade
- Combine com padrões de projeto

## ✨ Citação de Robert C. Martin (Uncle Bob)

📖 **Reflexão sobre Simplicidade no Código**

> 💡 *“Basta escrever código simples". 
>Este é um bom conselho.  
> No entanto, se os anos nos ensinaram alguma coisa, é que a simplicidade requer disciplinas guiadas por princípios.  
> São esses princípios que definem a simplicidade.  
> São essas disciplinas que obrigam os programadores a produzir código que se inclina para a simplicidade.*  
>
> 🚫 *A melhor maneira de complicar as coisas é dizer a todos para ‘serem simples’ e não dar mais nenhuma orientação.”*

---

👤 **Autor:** Robert C. Martin (*Uncle Bob*)  
📅 Publicado em 18 de outubro de 2020  
🔗 Fonte: [Solid Relevance – The Clean Coder Blog](https://blog.cleancoder.com/uncle-bob/2020/10/18/Solid-Relevance.html)


**🚀 Próximos Passos:**
- SOLID em GO
- SOLID aprofundado
- Padrões de projeto (Design Patterns)
- Arquitetura limpa (Clean Architecture)

**📖 Material de Apoio:**

- 🔗 [**Blog do Uncle Bob – Solid Relevance**](https://blog.cleancoder.com/uncle-bob/2020/10/18/Solid-Relevance.html)  
- 💻 [**Repositório no GitHub**](https://github.com/wellalencarweb/solid)  
- 📚 **Livro:** *Clean Code* – Robert C. Martin  
- 🐘 [**Documentação Oficial do PHP**](https://www.php.net)

