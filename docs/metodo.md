# Método

Os métodos Go são semelhantes à função Go com uma diferença, ou seja, o método contém um argumento de __receptor__ nele. Com a ajuda do argumento do receptor, o método pode acessar as propriedades do receptor. O receptor pode ser do tipo __struct ou do tipo non-struct__, qualquer tipo pode ter métodos, até mesmo um tipo de função ou tipos de alias para _int_, _bool_, _string_ ou _array_. Quando você cria um método em seu código, o receptor e o tipo de receptor devem estar presentes no __mesmo pacote__.

> __Note__
>
> - Você não tem permissão para criar um método no qual o tipo de receptor já está definido em outro pacote, incluindo tipos nativos como _int_, _string_, etc. Se você tentar fazer isso, o compilador dará um erro. Para usar os tipo nativos é necessário criar alias para estes tipos.
> - Os receptores são acessíveis dentro do método.

Abaixo o exemplo de sintaxe de método em Go:

```go
func (nome_receptor tipo_receptor) nome_funcao(lista_parametros)(tipo_retorno){
    // corpo do método.....
}
```

__Exemplo de método com receptor do tipo _struct_\__:

```go
package main

import "fmt"

type author struct {
    name      string
    branch    string
    particles int
    salary    int
}

func (a author) show() {
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.particles)
    fmt.Println("Salary: ", a.salary)
}

func main() {
    res := author{
        name: "Sona",
        branch: "CSE",
        particles: 203,
        salary: 34000,
    }
    res.show()
}

// Saída:
// Author's Name:  Sona
// Branch Name:  CSE
// Published articles:  203
// Salary:  34000
```

__Exemplo de método com receptor do tipo _non-struct_\__:

```go
package main
 
import "fmt"
 
type inteiro int
 
func (d1 inteiro) multiply(d2 inteiro) inteiro {
    return d1 * d2
}

func main() {
    value1 := inteiro(23)
    value2 := inteiro(20)
    res := value1.multiply(value2)
    fmt.Println("Final result: ", res)
}

// Saída: Final result:  460
```

No idioma Go, você tem permissão para criar um método com um receptor de ponteiro. Com a ajuda de um receptor de ponteiro, se uma alteração for feita no método, ela refletirá no chamador.

__Exemplo métodos com receptor de ponteiro__:

```go
package main

import "fmt"

type author   struct {
    name      string
    branch    string
    particles int
}

func (a *author) show(abranch string) {
    (*a).branch = abranch
}

func main() {

    res := author{
        name: "Sona",
        branch: "Local",
    }

    fmt.Println("Author's name: ", res.name)
    fmt.Println("Branch Name(Before): ", res.branch)

    p := &res
    p.show("Global")

    fmt.Println("Author's name: ", res.name)
    fmt.Println("Branch Name(After): ", res.branch)
}

// Saída:
// Author's name:  Sona
// Branch Name(Before):  Local
// Author's name:  Sona
// Branch Name(After):  Global
```

## Sobrecarga de método

A sobrecarga do método é possível com base no tipo de receptor, um método com o mesmo nome pode existir em dois tipos de receptores diferentes desde que estejam no mesmo pacote.

```go
package main

import "fmt"

type multiply int
type addition int

func (m multiply) twice() multiply {
    return multiply(m * 2)
}

func (a addition) twice() addition {
    return addition(a + a)
}

func main() {
    var mul multiply = 15
    tm := mul.twice()
    fmt.Println(tm)

    var add addition = 15
    ta := add.twice()
    fmt.Println(ta)
}

// Saída:
// 30
// 30
```

## Composição

Não há nenhum conceito de um tipo de classe que sirva de base para objetos no Go. Qualquer tipo de dados em Go pode ser usado como um objeto. Tipo _struct_ em Go pode receber um conjunto de métodos para definir seu comportamento. Não existe nenhum tipo especial chamado classe ou objeto no GO mas ele suporta a maioria dos conceitos que geralmente são atribuídos à programação orientada a objetos.

GO não suporta o polimorfismo através da herança. Mas é possível criar objetos e expressar suas relações polimórficas através da composição usando um tipo como uma _struct_ ou uma _interface_.

Abaixo um exemplo do uso de composição:

```go
package main

import "fmt"

type Animal struct{}

func (a Animal) Comer() {
    fmt.Println("Comendo")
}

type MembroFamilia struct{}

func (mf MembroFamilia) Nome() {
    fmt.Println("Meu nome não é Johnny")
}

type Cachorro struct {
    Animal        // Struct incorporada
    MembroFamilia // Struct incorporada
}

func main() {
    d := Cachorro{}
    d.Comer() // "Comendo"
    d.Nome()  // "Meu nome não é Johnny"
}
```

## Referências

- <https://www.geeksforgeeks.org/methods-in-golang>
- <https://www.golangprograms.com/golang/methods-interfaces-objects>