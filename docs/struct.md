# Struct

É possível programar orientado a objetos, mas não da forma mais comum, pois Go não utiliza classes e sim **estruturas**. Em Go são criados métodos sem classes, interface sem hierarquia, e reutilização de código sem herança (apenas composição).

Abaixo um exemplo de uma estrutura do tipo `Pessoa` contendo `Nome` com o tipo _string_ e `Idade` com o tipo _int_:

```go
package main

import "fmt"

type Pessoa struct {
    Nome string
    Idade int
}

func main() {
    c := Pessoa{ // instanciando e definindo os valores na struct
        Nome: "Sammy the Shark",
        Idade: 19,
    }
    fmt.Println(c.Nome)
}

// Saída: Sammy the Shark
```

Também podemos instanciar uma struct sem definir os valores. Neste caso serão definidos os valores _default_ de cada tipo do campo.

```go
package main

import "fmt"

type Pessoa struct {
    Nome string
    Idade int
}

func main() {
    c := Pessoa{} // ou: var c Pessoa

    fmt.Println(c.Nome)
    fmt.Println(c.Idade)
}

// Saída: 
// ""
// 0
```

## Visibilidade

Os campos de uma _struct_ seguem as mesmas regras de exportação que outros identificadores dentro da linguagem de programação Go:
    - Se um nome de campo começar com uma letra maiúscula, ele será legível e gravável por código fora do pacote onde a estrutura foi definida.
    - Se o campo começar com uma letra minúscula, somente o código dentro do pacote dessa _struct_ poderá ler e gravar esse campo.

```go
package main

type Pessoa struct {
    Nome string
    idade int // visível somente no pacote main
}
```

## Embedded Structs

Além de definir um novo tipo para representar uma _struct_, você também pode definir uma _struct_ incorporada. Essas definições de _struct_ em tempo real são úteis em situações em que inventar novos nomes para tipos de _struct_ seria um esforço desperdiçado.

```go
package main

import "fmt"

func main() {
    c := struct {
        Name string
        Type string
    }{
        Name: "Sammy",
        Type: "Shark",
    }
    fmt.Println(c.Name, "the", c.Type)
}

// Saída: Sammy the Shark
```

## Anonymous Embedded Fields

Um campo incorporado anônimo é um campo _struct_ que não tem um nome de campo explícito. As estruturas que têm um campo anônimo obtêm todos os métodos e propriedades da _struct_ aninhada. Esses métodos e propriedades são chamados de métodos e propriedades "promovidos". Campos incorporados anônimos também podem ser acessados diretamente pelo nome do tipo.

```go
package main

import "fmt"

type Residencia struct {
    Quartos int
}

type Predio struct {
    Residencia // campo anônimo
    Elevadores int
}

func main() {
    p := Predio{
        Residencia: Residencia{Quartos: 3},
        Elevadores:   1,
    }
    fmt.Println(p.Quartos) // campos de Residencia são incorporados
    fmt.Println(p.Elevadores)
}

// Saída:
// 3
// 1
```

## Referências

- <https://pt.wikipedia.org/wiki/Go_(linguagem_de_programa%C3%A7%C3%A3o)>
- <https://www.digitalocean.com/community/tutorials/defining-structs-in-go>
