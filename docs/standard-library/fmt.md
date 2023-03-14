# fmt

O pacote `fmt` implementa E/S formatada com funções análogas para C como `printf` e `scanf`. O formato dos 'verbos' é derivado de C, mas são mais simples.

## Verbos

### Geral

```txt
%v  - o valor em um formato padrão, ao imprimir structs, o sinalizador "+" (%+v) adiciona nomes de campo.
%#v - uma representação da sintaxe Go do valor.
%T  - uma representação da sintaxe Go do tipo do valor.
%%  - um sinal de porcentagem literal; não consome nenhum valor (imprime somente um %).

O formato padrão para %v é:

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x se impresso com %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

Exemplos:

```go
b := true
fmt.Printf("Bool: %v", b)       // Bool: true
fmt.Printf("Bool: %#v", b)      // Bool: true
fmt.Printf("Bool: %T", b)       // Bool: bool
i := 234
fmt.Printf("Int: %v", i)        // Int: 234
fmt.Printf("Int: %#v", i)       // Int: 234
fmt.Printf("Int: %T", i)        // Int: int
f := 12.34
fmt.Printf("Float: %v", f)      // Float: 12.34
fmt.Printf("Float: %#v", f)     // Float: 12.34
fmt.Printf("Float: %T", f)      // Float: float
p := 15
fmt.Printf("Pointer: %v", &p)  // Pointer: 0xc00001c038
fmt.Printf("Pointer: %#v", &p) // Pointer: (*int)(0xc00001c038)
fmt.Printf("Pointer: %T", &p)  // Pointer: *int
s := []int{1, 2, 3}
fmt.Printf("Slice: %v", s)     // Slice: [1 2 3]
fmt.Printf("Slice: %#v", s)    // Slice: []int{1, 2, 3}
fmt.Printf("Slice: %T", s)     // Slice: []int

type Teste struct {
    numero   int
    booleano bool
}
t := Teste{numero: 2, booleano: false}
fmt.Printf("Struct: %v", t)  // Struct: {2 false}
fmt.Printf("Struct: %#v", t) // Struct: main.Teste{numero:2, booleano:false}
fmt.Printf("Struct: %T", t)  // Struct: main.Teste
```

#### Outras flags

```txt
'+' - Sempre imprime um sinal para valores numéricos.
        - (%+q): garante saída somente ASCII.
'-' - Preenchimento com espaços à direita e não à esquerda (justifique o campo à esquerda)
'#' - Formato alternativo:
        - binário (%#b): adicione 0b inicial.
        - octal (%#o): adicione 0 inicial.
        - hex (%#x ou %#X): adicione 0x ou 0X inicial.
        - (%#p): suprime 0x para %p; 
        - (%q): imprime uma string raw (entre aspas) se strconv.CanBackquote retorna true.
        - (%e, %E, %f, %F, %g e %G): imprime sempre um ponto decimal.
        - (%g e %G): não remova zeros à direita
        - (%#U): escreve por exemplo U+0078 'x' se o caractere for imprimível.
' ' - Preenchimento com espaço:
        - (%d): deixa um espaço para o sinal omitido em números.
        - (% x, % X): coloca espaços entre bytes imprimindo strings ou slices em hexadecimal.
'0' - Preenchimento com zeros à esquerda em vez de espaços:
        - Números: move o preenchimento após o sinal.
        - Strings, bytes slices e byte arrays: será ignorado.
```

### Booleano

```txt
%t - a palavra verdadeiro ou falso
```

Exemplos:

```go
b := true
fmt.Printf("Bool: %t", b) // Bool: true
```

### Integer

```txt
%b - base 2.
%c - o caractere representado pelo ponto de código Unicode correspondente.
%d - base 10.
%o - base 8.
%O - base 8 com prefixo 0o.
%q - um caractere entre aspas simples escapado com segurança.
%x - base 16, com letras minúsculas para a-f
%X - base 16, com letras maiúsculas para A-F
%U - formato Unicode: U+1234; o mesmo que "U+%04X"

%4d  - Pad with spaces (tamanho 4, alinhado a direita)
%-4d - Pad with spaces (tamanho 4, alinhado a esquerda)
%04d - Pad with zeroes (tamanho 4)
```

Exemplos:

```go
i := 20
fmt.Printf("Int: %b", i) // Int: 10100
fmt.Printf("Int: %c", i) // Int: 
fmt.Printf("Int: %d", i) // Int: 20
fmt.Printf("Int: %o", i) // Int: 24
fmt.Printf("Int: %O", i) // Int: 0o24
fmt.Printf("Int: %q", i) // Int: '\x14'
fmt.Printf("Int: %x", i) // Int: 14
fmt.Printf("Int: %X", i) // Int: 14
fmt.Printf("Int: %U", i) // Int: U+0014

fmt.Printf("Int: %4d", i)  // Int:   20
fmt.Printf("Int: %-4d", i) // Int: 20  //eol
fmt.Printf("Int: %04d", i) // Int: 0020
```

### Ponto Flutuante

```txt
%b - notação científica sem decimal com expoente a potência de dois, na forma de strconv. FormatFloat com o formato 'b', por exemplo. -123456p-78.
%e - notação científica, por exemplo -1.234456e+78.
%E - notação científica, por exemplo -1.234456E+78.
%f - ponto decimal, mas sem expoente, ex. 123.456.
%F - sinônimo de %f.
%g %e - para grandes expoentes, %f caso contrário. A precisão é discutida abaixo.
%G %E - para grandes expoentes, %F caso contrário.
%x - notação hexadecimal (com potência decimal de dois expoentes), ex. -0x1.23abcp+20.
%X - notação hexadecimal maiúscula, por exemplo -0X1.23ABCP+20.

O tamanho é especificado por um número decimal opcional imediatamente anterior ao verbo. Se ausente, a quantidade será a que for necessário para representar o valor.

- Se o tamanho for maior que a quantidade de dígitos, o restante é preenchido com espaço.
- Se o tamanho for menor que a quantidade de dígitos, será usado o tamanho padrão.

A precisão (casas decimais) é especificada após o tamanho (opcional) e um ponto "." seguido por um número decimal. Se nenhum ponto "." estiver presente, uma precisão padrão será usada. Se houver ponto "." sem a precisão, será aplicado precisão zero.

%f tamanho padrão, precisão padrão
%9f tamanho 9, precisão padrão
%.2f tamanho padrão, precisão 2
%9.2f tamanho 9, precisão 2
%9.f tamanho 9, precisão 0

O tamanho e a precisão são medidas em unidades de pontos de código Unicode, ou seja, runes. 
```

Exemplos:

```go
f := 120.1234567890
fmt.Printf("Float: %b", f) // Float: 8452936800521817p-46
fmt.Printf("Float: %e", f) // Float: 1.201235e+02
fmt.Printf("Float: %E", f) // Float: 1.201235E+02
fmt.Printf("Float: %f", f) // Float: 120.123457
fmt.Printf("Float: %F", f) // Float: 120.123457
fmt.Printf("Float: %g", f) // Float: 120.123456789
fmt.Printf("Float: %G", f) // Float: 120.123456789
fmt.Printf("Float: %x", f) // Float: 0x1.e07e6b74dce59p+06
fmt.Printf("Float: %X", f) // Float: 0X1.E07E6B74DCE59P+06

fmt.Printf("Float: %9f", f)   // Float: 120.123457
fmt.Printf("Float: %.2f", f)  // Float: 120.12
fmt.Printf("Float: %9.2f", f) // Float:    120.12
fmt.Printf("Float: %1.1f", f) // Float: 120.1
fmt.Printf("Float: %6.1f", f) // Float:  120.1
fmt.Printf("Float: %9.f", f)  // Float:       120
fmt.Printf("Float: %.12f", f) // Float: 120.123456789000
```

### String e Slice de bytes

```txt
%s - os bytes não interpretados da string ou slice.
%q - uma string entre aspas duplas escapada com segurança.
%x - base 16, minúsculas, dois caracteres por byte.
%X - base 16, maiúsculas, dois caracteres por byte.
```

Exemplos:

```go
s := "abcdefghijklmnopqrstuvwxyz"
fmt.Printf("String: %s", s) // String: abcdefghijklmnopqrstuvwxyz
fmt.Printf("String: %q", s) // String: "abcdefghijklmnopqrstuvwxyz"
fmt.Printf("String: %x", s) // String: 6162636465666768696a6b6c6d6e6f707172737475767778797a
fmt.Printf("String: %X", s) // String: 6162636465666768696A6B6C6D6E6F707172737475767778797A

fmt.Printf("String: %30s", s)  // String:     abcdefghijklmnopqrstuvwxyz
fmt.Printf("String: %-30s", s) // String: abcdefghijklmnopqrstuvwxyz    //eol
fmt.Printf("String: %030s", s) // String: 0000abcdefghijklmnopqrstuvwxyz

bs := []byte(s)
fmt.Printf("String: %s", bs) // String: abcdefghijklmnopqrstuvwxyz
fmt.Printf("String: %d", bs) // String: [97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 121 122]

bs = []byte{97, 98, 99, 100, 101, 102, 103}
fmt.Printf("String: %s", bs) // String: abcdefg
```

### Slice

```txt
%p - endereço do 0º elemento na notação de base 16, com 0x inicial.
```

Exemplos:

```go
s := []int{1, 2, 3}
fmt.Printf("Slice: %p", s) // Slice: 0xc0000b2000
fmt.Printf("Slice: %d", s) // Slice: [1 2 3] // %d padrão de inteiros
```

### Ponteiro

```txt
%p - notação de base 16, com 0x à esquerda.
%b, %d, %o, %x e %X - também funcionam com ponteiros, formatando o valor exatamente como se fosse um número inteiro.
```

Exemplos:

```go
v := 5
p := &v
fmt.Printf("Pointer: %p", p) // Pointer: 0xc0000b2000
fmt.Printf("Pointer: %b", p) // Pointer: 1100000000000000000010110010000000000000
fmt.Printf("Pointer: %d", p) // Pointer: 824634449920
fmt.Printf("Pointer: %o", p) // Pointer: 14000000340060
fmt.Printf("Pointer: %x", p) // Pointer: c00010a000
fmt.Printf("Pointer: %X", p) // Pointer: C00010A000
```

### Objetos compostos

Para objetos compostos, os elementos são impressos usando essas regras, recursivamente:

```txt
struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
pointer to above:   &{}, &[], &map[]
```

Exemplos:

```go
type Teste struct {
    numero   int
    booleano bool
}
s := Teste{}
fmt.Printf("Struct: %v", s)  // Struct: {0 false}
fmt.Printf("Struct: %v", &s) // Struct: &{0 false}

m := make(map[string]int)
m["a"] = 1
m["b"] = 2
fmt.Printf("Map: %v", m)  // Map: map[a:1 b:2]
fmt.Printf("Map: %v", &m) // Map: &map[a:1 b:2]
```

## Funções

### Append

Formata usando os __formatos padrão__ para seus operandos, anexa o resultado ao _slice_ de byte e retornar o _slice_ atualizado.

```go
Append(b []byte, a ...any) []byte
```

Exemplos:

```go
b := []byte("abc")
s1 := "d"
s2 := "e"
s3 := "f"

r := fmt.Append(b, s1, s2, s3)
fmt.Printf("%v", r) // [97 98 99 100 101 102]
fmt.Printf("%s", r) // abcdef
```

### Appendln

Formata usando os __formatos padrão__ para seus operandos, anexa o resultado ao _slice_ de byte e retornar o _slice_ atualizado. É adicionado espaço entre os operandos e adicionado uma nova linha.

```go
Appendln(b []byte, a ...any) []byte
```

Exemplos:

```go
b := []byte("abc")
s1 := "d"
s2 := "e"
s3 := "f"

r := fmt.Appendln(b, s1, s2, s3)
fmt.Printf("%v", r) // [97 98 99 100 32 101 32 102 10]
fmt.Printf("%s", r) // abcd e f\n
```

### Appendf

Formata usando o __formatos especificado__, anexa o resultado ao _slice_ de byte e retornar o _slice_ atualizado.

```go
Appendf(b []byte, format string, a ...any) []byte
```

Exemplos:

```go
b := []byte("abc")
s1 := "d"
s2 := "e"
s3 := "f"

r := fmt.Appendf(b, "%s,%s,%s", s1, s2, s3)
fmt.Printf("%v", r) // [97 98 99 100 44 101 44 102]
fmt.Printf("%s", r) // abcd,e,f
```

### Errorf

Formata de acordo com um formato especificado e retorna a string como um valor do tipo `error`.

Se o formato especificado incluir um verbo `%w` com um operando de _error_, o _error_ retornado implementará um método __Unwrap__ retornando o operando. Se houver mais de um verbo `%w`, o _error_ retornado implementará um método _Unwrap_ retornando um `[]error` contendo todos os operandos `%w` na ordem em que aparecem nos argumentos. É inválido fornecer o verbo `%w` com um operando que não implemente a interface de _error_. Caso contrário, o verbo `%w` é sinônimo de `%v`.

```go
Errorf(format string, a ...any) error
```

Exemplos:

```go
const name, id = "João", 17
err := fmt.Errorf("user %q (id %d) not found", name, id)

fmt.Println(err.Error()) // user "João" (id 17) not found
```

```go
errN3 := fmt.Errorf("mensagem-%d", 3)
errN2 := fmt.Errorf("mensagem-%d%w", 2, errN3)
errN1 := fmt.Errorf("mensagem-%d%w", 1, errN2)

fmt.Println(errN1.Error())                       // mensagem-1mensagem-2mensagem-3
fmt.Println(errors.Unwrap(errN1))                // mensagem-2mensagem-3
fmt.Println(errors.Unwrap(errors.Unwrap(errN1))) // mensagem-3
```

### FormatString

Retorna uma string que representa a diretiva de formatação totalmente qualificada capturada pelo _fmt.State_, seguida pelo verbo do argumento. O próprio _state_ não contém o verbo. O resultado tem um sinal de porcentagem inicial seguido por quaisquer _flags_, a largura e a precisão. _Flags_, largura e precisão ausentes são omitidos. Esta função permite que um _Formatter_ reconstrua a diretiva original que aciona a chamada para _Format_.

```go
FormatString(state State, verb rune) string
```

### Print

Imprimir formatos usando os formatos padrão para seus operandos e gravar na saída padrão. Espaços são adicionados entre os operandos quando __nenhum__ deles é uma _string_. Ele retorna o número de _bytes_ gravados e qualquer erro de gravação encontrado.

```go
Print(a ...any) (n int, err error)
```

Exemplos:

```go
const name, age = "João", 22
fmt.Print(name, " tem ", age, " anos.\n") // João tem 22 anos.

// É convencional não se preocupar com nenhum erro retornado pelo Print.
```

### Println

Imprimir formatos usando os formatos padrão para seus operandos e gravar na saída padrão. Espaços são __sempre__ adicionados entre os operandos e uma __nova linha__ é acrescentada. Ele retorna o número de _bytes_ gravados e qualquer erro de gravação encontrado.

```go
Println(a ...any) (n int, err error)
```

Exemplos:

```go
const name, age = "João", 22
fmt.Println(name, " tem ", age, " anos.\n") // João tem 22 anos.\n

// É convencional não se preocupar com nenhum erro retornado pelo Print.
```

### Printf

Formata de acordo com um formato especificado e grava na saída padrão. Ele retorna o número de _bytes_ gravados e qualquer erro de gravação encontrado.

```go
    Printf(format string, a ...any) (n int, err error)
```

Exemplos:

```go
const name, age = "João", 22
fmt.Printf("%s tem %d anos.\n", name, age)

// É convencional não se preocupar com nenhum erro retornado pelo Print.
```

### Sprint

Formata usando os formatos padrão para seus operandos e __retorna a string resultante__. Espaços são adicionados entre os operandos quando nenhum deles é uma _string_.

```go
Sprint(a ...any) string
```

Exemplos:

```go
const name, age = "João", 22
s := fmt.Sprint(name, " tem ", age, " anos.\n")

fmt.Print(s) // João tem 22 anos.
```

### Sprintln

Formata usando os formatos padrão para seus operandos e retorna a string resultante. Espaços são sempre adicionados entre os operandos e uma __nova linha é acrescentada__.

```go
Sprintln(a ...any) string
```

Exemplos:

```go
const name, age = "João", 22
s := fmt.Sprintln(name, " tem ", age, " anos.\n")

fmt.Print(s) // João tem 22 anos.\n
```

### Sprintf

Formata de acordo com o formato especificado e retorna a string resultante.

```go
Sprintf(format string, a ...any) string
```

Exemplos:

```go
const name, age = "João", 22
s := fmt.Sprintf("%s tem %d anos.\n", name, age)

fmt.Print(s) // João tem 22 anos.
```

### Fprint

Formata usando os __formatos padrão__ para seus operandos e grava em `w`. Espaços são adicionados entre os operandos quando nenhum deles é uma string. Ele retorna o número de bytes escritos e qualquer erro de escrita encontrado.

```go
Fprint(w io.Writer, a ...any) (n int, err error)
```

Exemplos:

```go
const name, age = "João", 17
n, err := fmt.Fprint(os.Stdout, name, " tem ", age, " anos.\n")

if err != nil {
    //...
}
fmt.Print(n, " bytes escritos.\n")

// Saída:
// João tem 17 anos.
// 19 bytes escritos.
```

### Fprintln

Formata usando os formatos padrão para seus operandos e grava em `w`. Espaços são sempre adicionados entre os operandos e uma nova linha é acrescentada. Ele retorna o número de bytes gravados e qualquer erro de gravação encontrado.

```go
Fprintln(w io.Writer, a ...any) (n int, err error)
```

```go
const name, age = "João", 17
n, err := fmt.Fprintln(os.Stdout, name, "tem", age, "anos.")

if err != nil {
    // ...
}
fmt.Println(n, "bytes escritos.")

// Saída:
// João tem 17 anos.
// 19 bytes escritos.
```

### Fprintf

Formatas de acordo com o __formato especificado__ e grava em `w`. Ele retorna o número de bytes gravados e qualquer erro de gravação encontrado.

```go
Fprintf(w io.Writer, format string, a ...any) (n int, err error)
```

Exemplos:

```go
const name, age = "João", 17
n, err := fmt.Fprintf(os.Stdout, "%s tem %d anos.\n", name, age)

if err != nil {
    // ...
}
fmt.Printf("%d bytes escritos.\n", n)

// Saída:
// João tem 17 anos.
// 19 bytes escritos.
```

### Scan

Varre o texto lido da entrada padrão, armazenando valores sucessivos separados por espaços em argumentos sucessivos. Novas linhas contam como espaço. Ele retorna o número de itens verificados com sucesso. Se for menor que o número de argumentos, `err` informará o motivo.

```go
Scan(a ...any) (n int, err error)
```

Exemplos:

```go
var a string
var b int
var c float64

n, err := fmt.Scan(&a, &b, &c)
if err != nil {
    // ...
}
fmt.Printf("%d: %s, %d, %f", n, a, b, c)

// Entrada (alternativa 1):
// João <enter/nova linha>
// 22 <enter/nova linha>
// 1.74 <enter/nova linha>

// Entrada (alternativa 2):
// João 22 1.74

// Saída:
// 3: João, 22, 1.740000
```

### Scanln

Semelhante a `Scan`, mas ele __para de escanear__ em uma __nova linha__ e após o item final deve haver uma nova linha ou _EOF_. Portanto os valores devem ser informados apenas separados por espaço, uma vez que ao digitar `<enter>` irá adicionar uma nova linha e o escanear ocorrerá.

```go
Scanln(a ...any) (n int, err error)
```

### Scanf

Verifica o texto lido da entrada padrão, armazenando valores sucessivos separados por espaço em argumentos sucessivos, conforme determinado pelo formato. Ele retorna o número de itens verificados com sucesso. Se for menor que o número de argumentos, `err` informará o motivo. __As novas linhas na entrada devem corresponder às novas linhas no formato__. A única exceção: o verbo `%c` sempre verifica a próxima runa na entrada, mesmo que seja um espaço (ou tabulação, etc.) ou nova linha.

```go
Scanf(format string, a ...any) (n int, err error)
```

Exemplos:

```go
var a string
var b int
var c float64

n, err := fmt.Scanf("%s\n%d %f", &a, &b, &c)
if err != nil {
    // ...
}
fmt.Printf("%d: %s, %d, %f", n, a, b, c)

// Entrada:
// João
// 22 1.74
// Saída:
// 3: João, 22, 1.740000
```

### Sscan

Varre a __string do argumento__, armazenando valores sucessivos separados por espaços em argumentos sucessivos. Novas linhas contam como espaço. Ele retorna o número de itens verificados com sucesso. Se for menor que o número de argumentos, `err` informará o motivo.

```go
Sscan(str string, a ...any) (n int, err error)
```

Exemplos:

```go
var a string
var b int
var c float64

n, err := fmt.Sscan("João 22 1.74", &a, &b, &c)
if err != nil {
    // ...
}
fmt.Printf("%d: %s, %d, %f", n, a, b, c) // 3: João, 22, 1.740000
```

### Sscanln

Semelhante ao `Sscan`, mas para de escanear em uma nova linha e após o item final deve haver uma nova linha ou EOF.

```go
Sscanln(str string, a ...any) (n int, err error)
```

### Sscanf

Varre a string no argumento, armazenando valores sucessivos separados por espaços em argumentos sucessivos conforme determinado pelo formato. Ele retorna o número de itens analisados com sucesso. As novas linhas na entrada devem corresponder às novas linhas no formato.

```go
Sscanf(str string, format string, a ...any) (n int, err error)
```

Exemplos:

```go
var a string
var b int
var c float64

n, err := fmt.Sscanf("João 22 1.74", "%s %d %f", &a, &b, &c)
if err != nil {
    // ...
}
fmt.Printf("%d: %s, %d, %f", n, a, b, c) // 3: João, 22, 1.740000
```

### Fscan

Varre o texto lido de `r`, armazenando valores sucessivos separados por espaços em argumentos sucessivos. Novas linhas contam como espaço. Ele retorna o número de itens verificados com sucesso. Se for menor que o número de argumentos, `err` informará o motivo.

```go
    Fscan(r io.Reader, a ...any) (n int, err error)
```

Exemplos:

```go
var a1, a2 string
var b1, b2 int
var c1, c2 float64

s := `dmr 1771 1.61803398875
ken 271828 3.14159`
r := strings.NewReader(s)

n, err := fmt.Fscan(r, &a1, &b1, &c1, &a2, &b2, &c2)
if err != nil {
    // ...
}
fmt.Printf("%d: %s, %d, %f, %s, %d, %f\n", n, a1, b1, c1, a2, b2, c2)

// Saída:
// 6: dmr, 1771, 1.618034, ken, 271828, 3.141590
```

### Fscanln

Semelhante ao `Fscan`, mas para de escanear em uma nova linha e após o item final deve haver uma nova linha ou _EOF_.

```go
    Fscanln(r io.Reader, a ...any) (n int, err error)
```

Se fosse usado o mesmo exemplo de `Fscan` seria retornado o erro:

```bash
panic: unexpected newline`
```

Exemplos:

```go
var a string
var b int
var c float64

s := `dmr 1771 1.61803398875
ken 271828 3.14159`
r := strings.NewReader(s)

for {
    n, err := fmt.Fscanln(r, &a, &b, &c)
    if err == io.EOF {
        break
    }
    if err != nil {
        // ...
    }
    fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
}

// Saída:
// 3: dmr, 1771, 1.618034
// 3: ken, 271828, 3.141590
```

### Fscanf

Varre o texto lido de `r`, armazenando valores sucessivos separados por espaços em argumentos sucessivos, conforme determinado pelo formato. Ele retorna o número de itens analisados com sucesso. As novas linhas na entrada devem corresponder às novas linhas no formato.

```go
    Fscanf(r io.Reader, format string, a ...any) (n int, err error)
```

Exemplos:

```go
var (
    i int
    b bool
    s string
)
r := strings.NewReader("5 true gophers")
n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
if err != nil {
    // ...
}
fmt.Println(i, b, s)
fmt.Println(n)

// Saída:
// 5 true gophers
// 3
```

## Referências

- <https://pkg.go.dev/fmt>