# Variável

As variáveis em go são usadas para dar o nome a um local de memória para armazenar os dados de tipo específico.

__Declaração variável com inicialização__:

```go
var i int = 10
var s string = "Canada"
```

__Declaração variável sem inicialização__:

```go
var i int
var s string

i = 10
s = "Canada"
```

__Declaração com inferência de tipo__:

O tipo do valor atribuído à variável será usado como o tipo dessa variável.

```go
var i = 10
var s = "Canada"
```

__Declaração de múltiplas variáveis__:

```go
var fname, lname string = "John", "Doe"
m, n, o := 1, 2, 3
item, price := "Mobile", 2000
```

__Declaração curta de variáveis__:

O operador `:=` indica que a declaração de variável curta está sendo usada. Não há necessidade de usar a palavra-chave var ou declarar o tipo de variável.

```go
name := "John Doe"
```

## Escopo definido por colchetes

Golang usa escopo lexical baseado em blocos de código para determinar o escopo das variáveis. O bloco interno pode acessar suas variáveis definidas pelo bloco externo, mas o bloco externo não pode acessar as variáveis definidas pelo bloco interno.

```go
package main

import (
    "fmt"
)

var s = "Japan"

func main() {
    fmt.Println(s)
    x := true

    if x {
        y := 1
        if x != false {
            fmt.Println(s)
            fmt.Println(x)
            fmt.Println(y)
        }
    }
    fmt.Println(x)
    fmt.Println(y) // Erro!
}

// Saída:
// Japan
// Japan
// true
// 1
// true
```

## Visibilidade

- Se o nome de uma variável começar com uma **letra minúscula**, ela só poderá ser acessada **dentro do pacote atual**, isso é considerado como variáveis não exportadas.
- Se o nome de uma variável começar com uma **letra maiúscula**, ela poderá ser acessada de pacotes **dentro e fora do pacote atual**, isso é considerado como variáveis exportadas.

## Valores Zero

Se você declarar uma variável sem atribuir-lhe um valor, o Golang vinculará automaticamente um valor padrão (ou um valor zero) à variável.

```go
var quantity float32 // 0
var price int16 // 0
var product string // ""
var inStock bool // false
```

## Bloco de Declaração de Variável de Golang

A declaração de variáveis pode ser agrupada em blocos para maior legibilidade e qualidade de código.

```go
package main

import "fmt"

var (
    product  = "Mobile"
    quantity = 50
    price    = 50.50
    inStock  = true
)

func main() {
    fmt.Println(quantity)
    fmt.Println(price)
    fmt.Println(product)
    fmt.Println(inStock)
}
```

## Referências

- <https://www.golangprograms.com/go-language/variables.html>
- <https://www.includehelp.com/golang/variables-in-golang.aspx>