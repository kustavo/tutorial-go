# Introdução

A linguagem GoLang, também conhecida apenas como Go é uma linguagem de programação criada pela Google e lançada em código livre em novembro de 2009. É uma linguagem compilada e focada em produtividade e programação concorrente.

Grandes nomes da área da computação fizeram parte do desenvolvimento dessa linguagem, sendo eles, Robert Griesemer, Ken Thompson e Rob Pike. A linguagem Go foi desenvolvida visando solucionar os desafios de engenharia enfrentados pelos desenvolvedores e desenvolvedoras do Google ao utilizar a linguagem C.

Go tem como suas principais características:

- É uma linguagem compilada;
- Estaticamente e fortemente tipada, além de possuir o recurso de inferência de tipos ou duck typing;
- Possui um garbage collector integrado, prevenindo problemas de vazamento de memória e um gerenciamento de memória apropriado;
- É uma linguagem opinativa, ela segue um sistema de tipos delimitado e lança erros quando variáveis ou bibliotecas não utilizadas são declaradas no programa;
- É simples de compilar e empacotar, gerando binários que podem ser executados diretamente pelo sistema operacional sem a necessidade de instalar nenhum interpretador previamente;
- Extensa biblioteca padrão com ferramentas para comunicação HTTP, serialização e desserialização de dados, expressões regulares e muito mais.

Algumas funcionalidades ausentes em Go são tratamento de exceção, herança, asserção e sobrecarga de métodos. Os autores argumentam abertamente contra asserções e defendem a omissão de herança de tipos em favor da eficiência. Ao contrário de Java, vetores associativos são parte intrínseca da linguagem, assim como strings.

## Instalação (Linux)

1. Baixar a última versão em <https://go.dev/dl> e configurar as variáveis de ambiente para o local onde se encontra o diretório `go`. Geralmente o local `/usr/local/go` é o mais recomendado. Mais detalhes em <https://go.dev/doc/install>:

1. No arquivo `~/.profile` adicionar a linha:

    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```

1. Recarregar o arquivo `.profile`

    ```bash
    source ~/.profile
    ```

1. Verificar versão instalada:

    ```bash
    go version
    ```

## Palavras-chave

Palavras-chave ou palavras reservadas são as palavras em um idioma que são usadas para algum processo interno ou representam algumas ações predefinidas. Essas palavras não são, portanto, autorizadas a serem usadas como um identificador. Isso resultará em um erro de tempo de compilação.

Há um total de 25 palavras-chave presentes na linguagem Go até o momento (Go 1.20):

```txt
break     default      func    interface  select
case      defer        go      map        struct
chan      else         goto    package    switch
const     fallthrough  if      range      type
continue  for          import  return     var
```

- `const`, `func`, `import`, `package`, `type` e `var` são usados para declarar todos os tipos de elementos de código em programas Go.
- `chan`, `interface`, `map` e `struct` são usados como partes em algumas denotações de tipo composto.
- `break`, `case`, `continue`, `default`, `else`, `fallthrough`, `for`, `goto`, `if`, `range`, `return`, `select` e `switch` são usados para controlar o fluxo de código.
- `defer` e `go` também são palavras-chave de fluxo de controle, mas de outras maneiras específicas. Eles modificam as chamadas de função.

## Gerenciador de dependências

Go usa Módulos Go para configurar as dependências de pacotes para a importação de recursos. Os módulos Go são arquivos de configuração `go.mod` colocados no seu diretório de pacotes que dizem ao compilador de onde importar os pacotes. Além das dependências, é nesse arquivo onde o Go adiciona o nome do seu package e a versão Go utilizada.

```go
// arquivo: go.mod
module github.com/example/cmd

go 1.19

require gorm.io/gorm
replace github.com/example/logging => ../logging
```

A primeira linha desse arquivo diz ao compilador que o nosso pacote chama `cmd` e tem o caminho de arquivo `$GOPATH/src/` + `github.com/example/cmd`. A segunda linha diz a versão Go utilizada. A terceira linha diz ao compilador que o pacote `gorm` está localizado em `gorm.io/gorm`. E a quarta linha diz ao compilador que o pacote `github.com/example/logging` pode ser encontrado localmente em disco, no diretório `../logging`.

> __Note__
> Quando o pacote a ser usado é privado, devemos explicitar através do comando:
>
>```bash
> go env -w GOPRIVATE=<ulr-dominio>
> # Exemplos: 
> # go env -w GOPRIVATE=github.com/kustavo
> # go env -w GOPRIVATE=github.com/kustavo/*
>```

### Comando "go mod init"

Um arquivo `.mod` pode ser criado usando o comando abaixo:

```bash
go mod init <caminho-importacao>
# Exemplo: go mod init github.com/example
```

O que `go mod init` fará é criar o arquivo `go.mod` no diretório que será a raiz do módulo e descrever qual o caminho base, ou seja, ao importar nossos pacotes usaremos o caminho base `github.com/example`. Abaixo o contéudo do arquivo `go.mod`:

```go
// arquivo: go.mod
module github.com/example

go 1.19
```

### Comando "go get"

Para baixar um pacote e adicionar ao `go.mod` podemos usar o comando:

```bash
go get <caminho-pacote>
# Exemplo: go get github.com/gorilla/mux
```

O que `go get` fará neste caso é baixar o código fonte do GitHub e colocar os arquivos em `$GOPATH/src/github.com/gorilla/mux`.

Todos os pacotes são importados através de seu caminho completo começando de `$GOPATH/src`, o que explica a necessidade de definir o `$GOPATH` durante a instalação do Go. A única exceção para esta regra é a `stdib` que é importada de `$GOROOT/src`. Abaixo o conteúdo do arquivo `go.mod`:

```go
// arquivo: go.mod
module github.com/example

go 1.19

require github.com/gorilla/mux v1.5.2
```

### Comando "go install"

Esse comando compila todos os pacotes e gera os arquivos executáveis movendo-os para `$GOPATH/pkg` ou `$GOPATH/bin`.

```bash
go install
```

### Comando "go mod tidy"

Para atualizar o `go.mod`, ou seja, adicionar as dependências e remover dependências não usadas, podemos usar o comando:

```bash
go mod tidy
```

### Comando "go vet"

Examina o código-fonte do Go e relata construções suspeitas. Ele usa heurística que não garantem que todos os relatórios são problemas genuínos, mas pode encontrar erros não capturado pelos compiladores.

```bash
# Examina o pacote no diretório informado
go vet <path-pacote>
```

Para listar as verificações disponíveis, execute `go tool vet help`:

```txt
asmdecl      report mismatches between assembly files and Go declarations
assign       check for useless assignments
atomic       check for common mistakes using the sync/atomic package
bools        check for common mistakes involving boolean operators
buildtag     check that +build tags are well-formed and correctly located
cgocall      detect some violations of the cgo pointer passing rules
composites   check for unkeyed composite literals
copylocks    check for locks erroneously passed by value
httpresponse check for mistakes using HTTP responses
loopclosure  check references to loop variables from within nested functions
lostcancel   check cancel func returned by context.WithCancel is called
nilfunc      check for useless comparisons between functions and nil
printf       check consistency of Printf format strings and arguments
shift        check for shifts that equal or exceed the width of the integer
stdmethods   check signature of methods of well-known interfaces
structtag    check that struct field tags conform to reflect.StructTag.Get
tests        check for common mistaken usages of tests and examples
unmarshal    report passing non-pointer or non-interface values to unmarshal
unreachable  check for unreachable code
unsafeptr    check for invalid conversions of uintptr to unsafe.Pointer
unusedresult check for unused results of calls to some functions
```

### Comando "go fmt"

Aplica padrões de formatação de código ao seu código. Essas alterações de formatação não afetam a execução do código, em vez disso, melhoram a legibilidade da base de código, garantindo que o código seja visualmente consistente. concentra-se em coisas como recuo, espaço em branco, comentários e sucintidade geral do código.

```bash
go fmt <path-arquivo>
```

### Comando "golint"

É um linter mantido pelos desenvolvedores do Go. Destina-se a impor as convenções de codificação descritas em _Effective Go_ e _CodeReviewComments_. Essas mesmas convenções são usadas no projeto Go de código aberto e no Google. Para instalá-lo execute o comando `go get -u golang.org/x/lint/golint`.

[Veja aqui outros Linters muito úteis](https://github.com/golangci/awesome-go-linters)

### Checksum

O arquivo `go.sum` é responsável por manter todas as informações para checksum das dependências utilizadas no projeto.

Todos os packages que o Go referencia no `go.sum` foram adicionados ao seu diretório `$GOPATH/pkg/mod`. Com esse cache em mãos, o Go não precisa fazer download das dependências toda vez que você executar o projeto.

Esse cache é global, ou seja, serão compartilhados por todos os projetos que você tiver ou criar na sua máquina.

### Vendoring

Como podemos compartilhar um código e garantir que todos tenham as dependências baixas e, o mais importante, a versão correta de cada dependência? Isso pode ser feito através de __vendoring__, que basicamente permite que aplicações Go utilizem dependências não só de `$GOPATH/src`, mas também de um diretório chamado `vendor` dentro de cada projeto.

O compilador do Go primeiramente procurará pelos pacotes dentro do diretório `vendor`, antes de procurar em `$GOPATH`.

Você pode executar o comando abaixo para gerar uma nova pasta chamada `vendor`:

```bash
go mod vendor
```

### Workspaces

Os espaços de trabalho nos permite trabalhar com vários módulos (`go.mod`) simultaneamente. Cada módulo dentro de um espaço de trabalho é tratado como um módulo principal.

Anteriormente, para adicionar um recurso a um módulo e usá-lo em outro módulo, era necessário para publicar as alterações no primeiro módulo ou editar o arquivo `go.mod` do dependente com uma diretiva `replace` apontando para os módulo locais e não publicadas. E então reverter este apontamento depois da publicação destes módulos.

Para definir os workspaces podemos criar um arquivo `go.work` na raiz. Neste arquivo indicamos os caminho dos Workspaces (diretório do arquivo `go.mod`):

```bash
go work init <caminho-relativo-workspaces>
# Exemplo: go work init ./hello
```

Será criado o arquivo `go.work`:

```go
go 1.18

use ./hello // Caminho relativo para o Workspace
```

Para adicionar outros workspaces:

```bash
go work use <caminho-relativo-workspaces>
# Exemplo: go work init ./strutil
```

```go
go 1.18

use ./hello
use ./strutil
```

```txt
- hello/
    - main.go
    - go.mod
- strutil/
    - main.go
    - go.mod 
- go.work
- go.mod
```

## Servidor de documentação

Iniciando um servidor godoc da aplicação:

```bash
~/go/bin/godoc -http=:<porta>
# Exemplo: ~/go/bin/godoc -http=:6060
```

Abrir no navegador: `http://localhost:6060/pkg/<nome-modulo>`

## Referências

- <https://pt.wikipedia.org/wiki/Go_(linguagem_de_programa%C3%A7%C3%A3o)>
- <https://www.digitalocean.com/community/tutorials/understanding-package-visibility-in-go-pt>
- <https://www.treinaweb.com.br/blog/o-que-e-e-como-comecar-com-golang>
- <https://aprendagolang.com.br/2022/05/12/como-utilizar-go-workspaces/>
- <https://go101.org/article/keywords-and-identifiers.html>
