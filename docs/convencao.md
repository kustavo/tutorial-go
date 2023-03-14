# Convenções

Para escrever Go bem, é importante entender suas propriedades e expressões idiomáticas além de conhecer as convenções estabelecidas para programação no Go, como nomenclatura, formatação, programa construção, e assim por diante, para que os programas que você escreve ser fácil para outros programadores de Go entenderem.

Abaixo as principais referências oficiais para escrever um bom código em Go:

- <https://go.dev/doc/effective_go>
- <https://github.com/golang/go/wiki/CodeReviewComments>
- <https://google.github.io/styleguide/go/decisions>
- <https://go.dev/ref/spec>

## Formatação

Go adota uma abordagem incomum e deixa a máquina cuidar da maioria dos problemas de formatação. O programa `gofmt` lê um programa Go e emite o código fonte em um estilo padrão de recuo e alinhamento vertical, retendo e, se necessário, reformatando os comentários. Todo o código Go nos pacotes `Standard library` foram formatados com `gofmt`.

Alguns detalhes de formatação:

- __Recuo__: Usamos tabs para recuo e `gofmt` as coloca por padrão.
- __Comprimento da linha__: Não impõe limite de comprimento.
- __Parênteses__: Estruturas de controle (`if`, `for`, `switch`) não possuem parênteses em sua sintaxe.

## Comentários

Go fornece comentários de bloco `/* */` e comentários de linha `//`.

O programa (e servidor web) `godoc` processa os arquivos de origem Go para extrair a documentação sobre o conteúdo do pacote. Os comentários que aparecem antes das declarações _top-level_, sem novas linhas intermediárias, são extraídos junto com a declaração para servir como texto explicativo para o item. A natureza e o estilo desses comentários determinam a qualidade da documentação que o `godoc` produz.

Todo pacote deve ter um comentário de pacote, um comentário de bloco precedendo a cláusula do pacote. Para pacotes com vários arquivos, o comentário do pacote só precisa estar presente em um arquivo, e qualquer um servirá. O comentário do pacote deve apresentar o pacote e fornecer informações relevantes para o pacote como um todo. Ele aparecerá primeiro na página do `godoc` e deve configurar a documentação detalhada a seguir.

```go
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

Se o pacote for simples, o comentário do pacote pode ser breve.

```go
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

Os comentários não precisam de formatação extra, como banners de estrelas. A saída gerada pode nem ser apresentada em uma fonte de largura fixa, portanto, não dependa do espaçamento para alinhamento - `godoc`, como `gofmt`, cuida disso. Os comentários são texto simples não interpretado, portanto, HTML e outras anotações como _this_ serão reproduzidos literalmente e não devem ser usados. Um ajuste que o `godoc` faz é exibir o texto recuado em uma fonte de largura fixa, adequada para trechos de programa. O comentário do pacote para o pacote `fmt` usa isso com bons resultados.

Dentro de um pacote, qualquer comentário imediatamente anterior a uma declaração de nível superior serve como um comentário doc para essa declaração. Cada nome exportado (em maiúsculas) em um programa deve ter um comentário doc.

Os comentários de documentos funcionam melhor como frases completas, que permitem uma ampla variedade de apresentações automatizadas. A primeira frase deve ser um resumo de uma frase que começa com o nome que está sendo declarado.

```go
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
```

Se cada comentário doc começar com o nome do item que descreve, você pode usar o subcomando doc da ferramenta go e executar a saída por meio do `grep`. Imagine que você não conseguia se lembrar do nome "Compile", mas estava procurando a função de análise para expressões regulares, então você executou o comando,

```bash
go doc -all regexp | grep -i <palavra>
```

Se todos os comentários doc no pacote começassem com "Esta função...", o `grep` não o ajudaria a lembrar o nome. Mas como o pacote inicia cada comentário de documento com o nome, você veria algo assim, que lembra a palavra que você está procurando.

```bash
$ go doc -all regexp | grep -i parse

$ Compile **parse**s a regular expression and returns, if successful, a Regexp
MustCompile is like Compile but panics if the expression cannot be **parse**d.
**parse**d. It simplifies safe initialization of global variables holding
```

A sintaxe de declaração do Go permite o agrupamento de declarações. Um único comentário doc pode introduzir um grupo de constantes ou variáveis ​​relacionadas. Uma vez que toda a declaração é apresentada, tal comentário pode muitas vezes ser superficial.

```go
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)
```

O agrupamento também pode indicar relacionamentos entre itens, como o fato de que um conjunto de variáveis ​​ser protegido por um `mutex`.

```go
var (
    countLock   sync.Mutex
    inputCount  uint32
    outputCount uint32
    errorCount  uint32
)
```

## Nomenclatura

- Um nome de função não pode começar com um número.
- A visibilidade de um nome fora de um pacote é determinada pelo fato de seu primeiro caractere ser maiúsculo.
- Um nome deve começar com uma letra e pode ter qualquer número de letras e números adicionais.
- Se um nome consiste em várias palavras, cada palavra após a primeira deve ser maiúscula. Não use sublinhados (ex: empName, EmpAddress).
- Nomes são _case-sensitive_, ou seja, diferenciam maiúsculas de minúsculas (ex: carro, Carro e CARRO são três variáveis ​​diferentes).
- Para acrônimos como API, HTTP, etc ou nomes como ID e DB, convencionalmente, mantemos essas palavras em sua forma original. (ex: userID, productAPI)

### Nomenclatura de arquivos

Não há uma convenção documentada para nomes de arquivos Go (além de _\_test.go_, _\_goos_goarch_, etc.). Há preferência para nomes curtos e significativos, como _io.go_, _pipe.go_, ect. Não é comum ver nomes em _mixedCaps_. Para pacotes mais complicados com muitos arquivos, às vezes são usados sublinhados quando uma separação adicional é necessária. Veja exemplos:

#### snake_case

```yaml
marshaling:
- encoding/json/example_marshaling_test.go
- encoding/json/example_text_marshaling_test.go
- encoding/xml/example_marshaling_test.go
- encoding/xml/example_text_marshaling_test.go
mmap:
- cmd/compile/internal/gc/mapfile_mmap.go
- cmd/internal/bio/buf_mmap.go
- cmd/link/internal/ld/outbuf_mmap.go
- runtime/cgo_mmap.go
- runtime/export_mmap_test.go
- runtime/runtime_mmap_test.go
string:
- cmd/compile/internal/gc/class_string.go
- cmd/compile/internal/gc/op_string.go
- cmd/compile/internal/syntax/operator_string.go
- cmd/compile/internal/syntax/token_string.go
- cmd/compile/internal/types/etype_string.go
- cmd/internal/obj/abi_string.go
- cmd/internal/obj/addrtype_string.go
- cmd/internal/objabi/reloctype_string.go
- cmd/internal/objabi/symkind_string.go
- cmd/link/internal/sym/symkind_string.go
- debug/dwarf/attr_string.go
- debug/dwarf/class_string.go
- debug/dwarf/tag_string.go
- debug/macho/reloctype_string.go
- html/template/attr_string.go
- html/template/delim_string.go
- html/template/element_string.go
- html/template/jsctx_string.go
- html/template/state_string.go
- html/template/urlpart_string.go
- math/big/accuracy_string.go
- math/big/roundingmode_string.go
- regexp/syntax/op_string.go
sysnum:
- internal/syscall/unix/at_sysnum_darwin.go
- internal/syscall/unix/at_sysnum_dragonfly.go
- internal/syscall/unix/at_sysnum_fstatat64_linux.go
- internal/syscall/unix/at_sysnum_fstatat_linux.go
- internal/syscall/unix/at_sysnum_linux.go
- internal/syscall/unix/at_sysnum_netbsd.go
- internal/syscall/unix/at_sysnum_newfstatat_linux.go
- internal/syscall/unix/at_sysnum_openbsd.go
```

#### lowercase

```yaml
gccgoinstallation:
- go/internal/gccgoimporter/gccgoinstallation_test.go
loopreschedchecks:
- cmd/compile/internal/ssa/loopreschedchecks.go
mkfastlog2table:
- runtime/mkfastlog2table.go
mksizeclasses:
- runtime/mksizeclasses.go
obscuretestdata:
- internal/obscuretestdata/obscuretestdata.go
reproduciblebuilds:
- cmd/compile/internal/gc/reproduciblebuilds_test.go
```

#### mixedCaps

```yaml
386Ops:
- cmd/compile/internal/ssa/gen/386Ops.go
"387":
- cmd/compile/internal/x86/387.go
AMD64Ops:
- cmd/compile/internal/ssa/gen/AMD64Ops.go
ARM64Ops:
- cmd/compile/internal/ssa/gen/ARM64Ops.go
ARMOps:
- cmd/compile/internal/ssa/gen/ARMOps.go
MIPS64Ops:
- cmd/compile/internal/ssa/gen/MIPS64Ops.go
MIPSOps:
- cmd/compile/internal/ssa/gen/MIPSOps.go
PPC64Ops:
- cmd/compile/internal/ssa/gen/PPC64Ops.go
S390XOps:
- cmd/compile/internal/ssa/gen/S390XOps.go
WasmOps:
- cmd/compile/internal/ssa/gen/WasmOps.go
arithBoundary:
- cmd/compile/internal/gc/testdata/arithBoundary_test.go
arithBoundaryGen:
- cmd/compile/internal/gc/testdata/gen/arithBoundaryGen.go
arithConst:
- cmd/compile/internal/gc/testdata/arithConst_test.go
arithConstGen:
- cmd/compile/internal/gc/testdata/gen/arithConstGen.go
deferNoReturn:
- cmd/compile/internal/gc/testdata/deferNoReturn_test.go
```

### Nomenclatura de pacotes

Por convenção, os pacotes recebem nomes com letras minúsculas e uma única palavra no singular; não deve haver necessidade de sublinhados ou _mixedCaps_. Exemplo:

```txt
time (provides functionality for measuring and displaying time)
list (implements a doubly linked list)
http (provides HTTP client and server implementations)
```

Erre no lado da brevidade, já que todos que usarem seu pacote estarão digitando esse nome. E não se preocupe com colisões a priori. O nome do pacote é apenas o nome padrão para importações; ele não precisa ser exclusivo em todo o código-fonte e, no caso raro de uma colisão, o pacote importador pode escolher um nome diferente para usar localmente.

Outra convenção é que o nome do pacote é o nome base de seu diretório de origem; o pacote em `src/encoding/base64` é importado como `encoding/base64`, mas tem o nome `base64`. Exemplo:

```txt
time -> src/time/
list -> src/container/list/
http -> src/net/http/
```

### Nomenclatura de interface

Por convenção, as interfaces de um método são nomeadas pelo nome do método mais um sufixo `-er` ou modificação semelhante para construir um substantivo de agente: _Reader_, _Writer_, _Formatter_, _CloseNotifier_ etc.

> __Note__
> A regra geral é `NomeMetodo + er = NomeInterface`. A parte complicada aqui é quando você tem uma interface com mais de um método. Nomear seguindo a convenção nem sempre será óbvio. Devo dividir a interface em várias interfaces com um único método? Acho que é uma decisão subjetiva que depende de cada caso.

### Nomenclatura de Getters e Setters

Go não oferece suporte automático para `getters` e `setters`. Não há nada de errado você mesmo fornecer `getters` e `setters`, e muitas vezes é apropriado fazê-lo, mas não é idiomático nem necessário colocar `Get` no nome do `getter`. Se você tem um campo chamado "owner" (minúsculas, não exportado), o método `getter` deve ser chamado "Owner" (maiúsculas, exportado), não "GetOwner". O uso de nomes em maiúsculas para exportação fornece o meio para discriminar o campo do método. Uma função `setter`, se necessário, provavelmente será chamada "SetOwner". Ambos os nomes são bem lidos na prática:

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

## Pacotes especiais

### Pacote main

Todo programa executável deve conter um pacote chamado `main` e uma função chamada `main`. Depois que seu programa for compilado e, quando você quiser executá-lo, a função `main` deste pacote será a primeira função a ser chamada.

> __Note__
> Se a função `init` estiver presente no mesmo arquivo, ela será executada antes de `main`. A função `init` pode conter todas as tarefas de inicialização necessárias para que o programa seja executado corretamente.

## Diretórios especiais

### Diretório internal

Colocar um pacote dentro de um diretório chamado `internal`, esconde ainda mais as estruturas internas do pacote, traz mais encapsulamento. Este pacote só poderá ser importado dos pacotes de seu diretório pai.

O diretório `internal` é usado para tornar pacotes específicos não importáveis.

[Veja aqui mais informações](package#internal-packages)

### Diretório vendoring

Como podemos compartilhar um código e garantir que todos tenham as dependências baixas e, o mais importante, a versão correta de cada dependência? Isso pode ser feito através de __vendoring__, que é basicamente á uma funcionalidade que permite aplicações Go utilizar dependências não só de `$GOPATH/src`, mas também de um diretório chamado `vendor` dentro de cada projeto. O compilador do Go primeiramente procurará pelos pacotes dentro do diretório `vendor`, antes de procurar em `$GOPATH`.

[Veja aqui mais informações](package#vendoring-packages)

## Organização de projeto

Uma proposta de Organização de projeto bem aceita pela comunidade Go contém os seguintes diretórios:

### Diretórios Go

#### /cmd

Este diretório contém os principais arquivos de ponto de entrada do aplicativo para o projeto. O nome do diretório para cada aplicação deve corresponder ao nome do executável (binário) que você deseja ter (ex. `/cmd/myapp/`). Arquivos deste diretório definem um pacote `main`, portanto o método `main` também estará num desses arquivos.

É comum ter uma pequena função `main` que importa e invoca o código dos diretórios `/internal` e `/pkg` e nada mais.

[Veja aqui alguns exemplos](https://github.com/golang-standards/project-layout/blob/master/cmd/README.md)

#### /internal

Este diretório contém aplicação privada e código de bibliotecas. Este é o código que você não quer que outras pessoas importem em suas aplicações ou bibliotecas.

[Veja aqui mais informações](package#internal-packages)

#### /vendor

Este diretório contém as dependências de aplicativos (gerenciadas manualmente). O comando `go mod vendor` criará o diretório `/vendor` para você.

[Veja aqui mais informações](package#vendoring-packages)

#### /pkg

Este diretório contém código de bibliotecas que podem ser usados por aplicativos externos (ex. `/pkg/mypubliclib`). Outros projetos irão importar essas bibliotecas esperando que funcionem, então pense duas vezes antes de colocar algo aqui.

É também uma forma de agrupar o código Go em um só lugar quando o diretório raiz contém muitos componentes e diretórios não Go, tornando mais fácil executar várias ferramentas Go.

Este é um padrão de layout comum, mas __não é universalmente aceito__ e alguns na comunidade Go não o recomendam.

Não há problema em não usá-lo se o projeto do seu aplicativo for muito pequeno e onde um nível extra de aninhamento não agrega muito valor. Pense nisso quando estiver ficando grande o suficiente e seu diretório raiz ficar muito ocupado (especialmente se você tiver muitos componentes de aplicativos não Go).

[Veja aqui alguns exemplos](https://github.com/golang-standards/project-layout/blob/master/pkg/README.md)

### Diretórios de aplicativos de serviço

#### /api

Este diretório contém especificações _OpenAPI/Swagger_, arquivos de esquema JSON, arquivos de definição de protocolo.

[Veja aqui alguns exemplos](https://github.com/golang-standards/project-layout/blob/master/api/README.md)

### Diretórios de aplicativos da web

#### /web

Este diretório contém componentes específicos de aplicativos da Web: ativos estáticos da Web, modelos do lado do servidor e SPAs.

### Diretórios de aplicativos comuns

#### /configs

Este diretório contém modelos de arquivo de configuração ou configurações padrão. Coloque seus arquivos de modelo _confd_ ou _consul-template_ aqui.

#### /init

Este diretório contém configurações de inicialização do sistema (_systemd_, _upstart_, _sysv_) e gerenciador/supervisor de processos (_runit_, _supervisord_).

#### /scripts

Este diretório contém scripts para executar várias operações de construção, instalação, análise, etc.

[Veja aqui alguns exemplos](https://github.com/golang-standards/project-layout/blob/master/scripts/README.md)

#### /build

Este diretório contém arquivos de empacotamento e integração contínua.

Coloque suas configurações de pacote e scripts em nuvem (_AMI_), contêiner (_Docker_), sistema operacional (_deb_, _rpm_, _pkg_) no diretório `/build/package`.

Coloque suas configurações e scripts de CI (travis, circle, drone) no diretório `/build/ci`. Observe que algumas das ferramentas de CI (por exemplo, Travis CI) são muito exigentes quanto à localização de seus arquivos de configuração. Tente colocar os arquivos de configuração no diretório `/build/ci` vinculando-os ao local onde as ferramentas de CI os esperam (quando possível).

#### /deployments

Este diretório contém arquivos de IaaS, PaaS, configurações e modelos de implantação de orquestração de sistema e contêiner (_docker-compose_, _kubernetes_ / _helm_, _mesos_, _terraform_, _bosh_).

> __Note__
> Observe que em alguns repositórios (especialmente em aplicativos implantados com kubernetes), esse diretório é denominado `/deploy`.

#### /test

Este diretório contém aplicações de testes externos adicionais e dados de teste. Sinta-se à vontade para estruturar o diretório `/test` da maneira que quiser. Para projetos maiores, faz sentido ter um subdiretório de dados. Por exemplo, você pode ter `/test/data` ou `/test/testdata` se precisar que o Go ignore o que está naquele diretório. Observe que o Go também irá ignorar diretórios ou arquivos que começam com `.` ou `_`, para que você tenha mais flexibilidade em termos de como nomear seu diretório de dados de teste.

[Veja aqui alguns exemplos](https://github.com/golang-standards/project-layout/blob/master/test/README.md)

### Outros diretórios

#### /docs

Este diretório contém documentos do projeto e do usuário (além da documentação gerada pelo godoc).

[Veja aqui alguns exemplos](https://github.com/golang-standards/project-layout/blob/master/docs/README.md)

#### /tools

Este diretório contém ferramentas de suporte para este projeto. Observe que essas ferramentas podem importar código dos diretórios `/pkg` e `/internal`.

[Veja aqui alguns exemplos](https://github.com/golang-standards/project-layout/blob/master/tools/README.md)

#### /examples

Este diretório contém exemplos para seus aplicativos e/ou bibliotecas públicas.

[Veja aqui alguns exemplos](https://github.com/golang-standards/project-layout/blob/master/examples/README.md)

#### /third_party

Este diretório contém ferramentas auxiliares externas, código bifurcado e outros utilitários de terceiros (por exemplo, Swagger UI).

#### /githooks

Este diretório contém Git hooks.

#### /assets

Este diretório contém outros recursos para acompanhar seu repositório (imagens, logotipos etc).

#### /website

Este diretório contém os dados do site do seu projeto se você não estiver usando as páginas do GitHub.

[Veja aqui alguns exemplos](https://github.com/golang-standards/project-layout/blob/master/website/README.md)

## Referências

- <https://go.dev/doc/effective_go>
- <https://www.golangprograms.com/naming-conventions-for-golang-functions.html>
- <https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md>
