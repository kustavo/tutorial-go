# Função

Uma função é um grupo de instruções que existem dentro de um programa com a finalidade de executar uma tarefa específica. Em um alto nível, uma função recebe uma entrada e retorna uma saída.

Funções são geralmente o bloco de códigos ou instruções em um programa que dá ao usuário a capacidade de reutilizar o mesmo código que, em última análise, economiza o uso excessivo de memória, atua como uma economia de tempo e, mais importante, fornece melhor legibilidade do código. Então, basicamente, uma função é uma coleção de instruções que executam alguma tarefa específica e retornam o resultado para o chamador. Uma função também pode executar alguma tarefa específica sem retornar nada.

Abaixo o exemplo de sintaxe de função em Go:

```go
func nome_funcao(lista_parametros)(tipo_retorno){
    // corpo da função.....
}
```

A declaração da função contém:

- __func__: É uma palavra-chave na linguagem Go, que é usada para criar uma função.
- __nome_funcao__: É o nome da função.
- __lista_parametros__: Ele contém o nome e o tipo dos parâmetros de função.
- __tipo_retorno__: É opcional e contém os tipos dos valores que a função retorna. Se você estiver usando __tipo_retorno__ em sua função, então é necessário usar uma instrução `return` em sua função.

A Invocação de Função ou a Chamada de Função é feita quando o usuário deseja executar a função. A função precisa ser chamada para usar sua funcionalidade. Como mostrado no exemplo abaixo, temos uma função chamada `area()` com dois parâmetros. Agora chamamos essa função na função principal usando seu nome, ou seja, `area(12, 10)` com dois parâmetros.

```go
package main
import "fmt"
 
func area(length, width int)int{
     
    Ar := length * width
    return Ar
}

func main() {
   fmt.Printf("Area of rectangle is : %d", area(12, 10))
}

// Saída:
// Area of rectangle is : 120
```

> __Note__
> Se as funções com nomes que começam com uma letra maiúscula serão exportadas para outros pacotes. Se o nome da função começar com uma letra minúscula, ela não será exportada para outros pacotes, mas você poderá chamar essa função dentro do mesmo pacote.

## Retornos

Golang permite que você nomeie os valores de retorno de uma função. Também podemos nomear o valor de retorno definindo variáveis. Neste caso não é necessário retornar a variável declarada no retorno. Veja o exemplo abaixo:

```go
package main

import "fmt"

func rectangle(l int, b int) (area int) {
    var parameter int
    parameter = 2 * (l + b)
    fmt.Println("Parameter: ", parameter)

    area = l * b
    return // Return statement without specify variable name
}

func main() {
    fmt.Println("Area: ", rectangle(20, 30))
}

// Saída:
// Parameter:  100
// Area:  600
```

## Argumentos

Na linguagem Go, os parâmetros passados para uma função são chamados de __parâmetros reais__, enquanto os parâmetros recebidos por uma função são chamados de __parâmetros formais__.

A linguagem Go oferece suporte a duas maneiras de passar argumentos para sua função:

- __Chamada por valor__: Neste método de passagem de parâmetros, os valores dos parâmetros reais são copiados para os parâmetros formais da função e os dois tipos de parâmetros são armazenados em diferentes locais de memória. Portanto, quaisquer alterações feitas dentro das funções __não são refletidas nos parâmetros reais__ do chamador.

    Exemplo:

    ```go
    package main
    
    import "fmt"
    
    func swap(a, b int) int {
        var o int
        o = a
        a = b
        b = o
        return o
    }
    
    func main() {
        var p int = 10
        var q int = 20
        
        fmt.Printf("p = %d and q = %d", p, q)
        swap(p, q) // chamada por valor
        fmt.Printf("\np = %d and q = %d",p, q)
    }

    // Saída:
    // p = 10 and q = 20
    // p = 10 and q = 20
    ```

- __Chamada por referência__: Ambos os parâmetros reais e formais referem-se aos mesmos locais, de modo que quaisquer alterações feitas dentro da função __são realmente refletidas nos parâmetros reais__ do chamador. Isto é realizado utilizando ponteiros, que passa o endereço de um tipo para a função. A pilha de funções tem uma referência ao objeto original. Portanto, quaisquer modificações no objeto passado modificarão o objeto original.

    Exemplo:

    ```go
    package main
    
    import "fmt"
    
    func swap(a, b *int) int {
        var o int
        o = *a
        *a = *b
        *b = o
        return o
    }
    
    func main() {
        var p int = 10
        var q int = 20

        fmt.Printf("p = %d and q = %d", p, q)
        swap(&p, &q) // chamada por referência
        fmt.Printf("\np = %d and q = %d",p, q)
    }

    // Saída:
    // p = 10 and q = 20
    // p = 20 and q = 10
    ```

## Função variádica

É a função que é chamada com o número variável de argumentos. Um usuário tem permissão para passar zero ou mais argumentos na função variádica. `Fmt.Printf` é um exemplo de função variádica, ele exigiu um argumento fixo no início mas depois ele pode aceitar qualquer número de argumentos.

Na declaração da função variádica, o tipo do __último parâmetro__ é precedido por uma reticência, ou seja, (`...`). Ele indica que a função pode ser chamada em qualquer número de parâmetros desse tipo.

Sintaxe:  

```go
function nome_funcao(param1, param2 ...tipo) tipo_retorno {
    // corpo da função.....
}
```

> __Note__
>
> - O parâmetro de função variádica deve ser sempre o __último__ e pode haver somente __um__ parâmetro de função variádica.
> - `...tipo` se comporta como um _slice_. Por exemplo, suponha que temos uma assinatura de função, ou seja, `add(b... int) int`, agora o parâmetro é do tipo `[]int`.
> - Você pode passar um _slice_ existente em uma função variádica.
> - Quando você não passa nenhum argumento na função variádica, então o _slice_ dentro da função será vazio.
> - As funções variádicas são geralmente usadas para funções que executam a formatação de cadeia de caracteres.
> - Você pode passar várias slice na função variádica.
> - Você __não__ pode usar parâmetro variádico como um valor de retorno, mas você pode retorná-lo como um _slice_.

__Exemplo com somente parâmetros de função variádica__:

```go
package main

import "fmt"

func main() {
    variadicExample("red", "blue", "green", "yellow")
}

func variadicExample(s ...string) {
    fmt.Println(s[0])
    fmt.Println(s[3])
}

// Saída:
// red
// yellow
```

__Exemplo com parâmetro de função normal com parâmetro de função variádica__:

```go
package main

import "fmt"

func main() {
    fmt.Println(calculation("Rectangle", 20, 30))
    fmt.Println(calculation("Square", 20))
}

func calculation(str string, y ...int) int {

    area := 1

    for _, val := range y {
        if str == "Rectangle" {
            area *= val
        } else if str == "Square" {
            area = val * val
        }
    }
    return area
}

// Saída
// 600
// 400
```

__Exemplo com diferentes tipos de argumentos em função variádica__:

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    variadicExample(1, "red", true, 10.5, []string{"foo", "bar", "baz"}, map[string]int{"apple": 23, "tomato": 13})
}

func variadicExample(i ...interface{}) {
    for _, v := range i {
        fmt.Println(v, "--", reflect.ValueOf(v).Kind())
    }
}

// Saída:
// 1 -- int
// red -- string
// true -- bool
// 10.5 -- float64
// [foo bar baz] -- slice
// map[apple:23 tomato:13] -- map
```

## Função anônima

Uma função anônima é uma função que foi declarada sem qualquer identificador nomeado para se referir a ela. Funções anônimas podem aceitar entradas e retornar saídas, assim como as funções padrão fazem.

__Exemplo atribuindo função à variável__:

```go
package main

import "fmt"

var (
    area = func(l int, b int) int {
        return l * b
    }
)

func main() {
    fmt.Println(area(20, 30))
}

// Saída: 600
```

__Exemplo passando argumentos__:

```go
package main

import "fmt"

func main() {
    func(l int, b int) {
        fmt.Println(l * b)
    }(20, 30)
}

// Saída: 600
```

__Exemplo passando argumentos e retornando valor__:

```go
package main

import "fmt"

func main() {
    fmt.Printf(
        "100 (°F) = %.2f (°C)\n",
        func(f float64) float64 {
            return (f - 32.0) * (5.0 / 9.0)
        }(100),
    )
}

// Saída: 100 (°F) = 37.78 (°C)
```

## Função de ordem superior

Uma função de Ordem Superior é uma função que recebe uma função como um argumento ou retorna uma função como saída.

__Exemplo passando uma função como argumento__:

```go
package main

import "fmt"

func sum(x, y int) int {
    return x + y
}
func partialSum(x int) func(int) int {
    return func(y int) int {
        return sum(x, y)
    }
}
func main() {
    partial := partialSum(3)
    fmt.Println(partial(7))
}

// Saída: 10
```

__Exemplo retornando funções de outras funções__:

```go
package main

import "fmt"

func squareSum(x int) func(int) func(int) int {
    return func(y int) func(int) int {
        return func(z int) int {
            return x*x + y*y + z*z
        }
    }
}
func main() {
    // 5*5 + 6*6 + 7*7
    fmt.Println(squareSum(5)(6)(7))
}

// Saída: 110
```

## Tipos de função definidos pelo usuário

Golang também suporta a definição de nossos próprios tipos de função.

```go
package main

import "fmt"

type First func(int) int
type Second func(int) First

func squareSum(x int) Second {
    return func(y int) First {
        return func(z int) int {
            return x*x + y*y + z*z
        }
    }
}

func main() {
    // 5*5 + 6*6 + 7*7
    fmt.Println(squareSum(5)(6)(7))
}

// Saída: 110
```

## Identificador em Branco

O `_` (sublinhado) em Golang é conhecido como o Identificador em Branco. Identificadores são o nome definido pelo usuário dos componentes do programa usados para a finalidade de identificação. O Golang tem um recurso especial para definir e usar a variável não utilizada usando o Identificador em Branco. Variáveis não utilizadas são aquelas variáveis que são definidas pelo usuário ao longo do programa mas que nunca são usadas.

O uso real do Identificador em Branco vem quando uma função retorna vários valores, mas precisamos apenas de alguns valores e queremos descartar alguns valores. Basicamente, ele diz ao compilador que essa variável não é necessária e pode ignorá-la sem qualquer erro. Veja o exemplo abaixo:

```go
package main
 
import "fmt"
 
func main() {
    mul, _ := mul_div(105, 7) // obtendo valor da multiplicação e ignorando o da divisão
    fmt.Println("105 x 7 = ", mul)
}

func mul_div(n1 int, n2 int) (int, int) {
     return n1 * n2, n1 / n2
}
```

## Chamas de funções diferidas

Go tem uma instrução especial chamada `defer` que programa uma chamada de função para ser executada após a conclusão da função. Considere o seguinte exemplo:

```go
package main

import "fmt"

func first() {
    fmt.Println("First")
}
func second() {
    fmt.Println("Second")
}
func main() {
    defer second()
    first()
}

// Saída:
// First
// Second
```

Uma instrução de diferimento (adiamento) é frequentemente usada com operações emparelhadas, como abrir e fechar, conectar e desconectar ou bloquear e desbloquear para garantir que os recursos sejam liberados em todos os casos, não importa o quão complexo seja o fluxo de controle. O lugar certo para uma instrução de adiamento que libera um recurso é __imediatamente após o recurso ter sido adquirido com êxito__.

Abaixo está o exemplo para abrir um arquivo e executar a ação de leitura/gravação sem usar o `defer`:

```go
func ReadWrite() bool {
    file.Open("file")

    if failureX {
        file.Close()
        return false
    }
    if failureY {
        file.Close()
        return false
    }
    file.Close()
    return true
}
```

Mesmo exemplo mas agora usando `defer`:

```go
func ReadWrite() bool {
    file.Open("file")
    defer file.Close()

    if failureX {
        return false
    }
    if failureY {
        return false
    }
    return true
}
```

> __Note__
>
> - Pode haver mais de uma função diferida no mesmo escopo.
> - As funções diferidas são executadas em ordem LIFO(Last-In, First-Out), ou seja, em forma de pilha.
> - Nas instruções `defer`, os argumentos são avaliados quando a instrução `defer` é executada, não quando é chamada.
> - As Funções diferidas são executadas mesmo se ocorrer um __pânico__ (em inglês, _panic_) de tempo de execução.

## Referências

- <https://www.geeksforgeeks.org/functions-in-go-language>
- <https://www.geeksforgeeks.org/variadic-functions-in-go>
- <https://www.golangprograms.com/go-language/variadic-functions.html>
- <https://www.golangprograms.com/go-language/functions.html>
- <https://www.golangprograms.com/go-language/deferred-functions-calls.html>
- <https://www.geeksforgeeks.org/defer-keyword-in-golang>