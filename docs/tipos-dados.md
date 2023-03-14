# Tipos de dados

Os tipos de dados especificam o tipo de dados que uma variável Go pode conter. Podem ser divididos em quatro categorias que são as seguintes:

- __Tipo básico__: Números, _strings_ e booleanos.
- __Tipo agregado__: _Array_ e _structs_.
- __Tipo de referência__: Ponteiros, _slices_, _maps_, funções e _channels_.
- __Tipo de interface__: Interfaces.

Os tipos numéricos ainda podem ser classificados nos seguintes tipos:

| Tipo    | Descrição |
|---------|-----------|
| int8    | Inteiro de 8 bits (-128 até 127). |
| int16   | Inteiro de 16 bits (-32768 até 32767). |
| int32   | Inteiro de 32 bits (-2147483648 até 2147483647). |
| int64   | Inteiro de 64 bits (-9223372036854775808 até 9223372036854775807). |
| uint8   | Inteiro sem sinal de 8 bits (0 até 255).|
| uint16  | Inteiro sem sinal de 16 bits (0 até 65535). |
| uint32  | Inteiro sem sinal de 32 bits (0 até 4294967295). |
| uint64  | Inteiro sem sinal de 64 bits (0 até 18446744073709551615). |
| int     | Inteiro de 32 ou 64 bits. |
| uint    | Inteiro sem sinal de 32 ou 64 bits. |
| rune    | É um sinônimo de int32 e também representa pontos de código Unicode. |
| byte    | É sinônimo de uint8. |
| uintptr | É um tipo inteiro sem sinal. Sua largura não é definida, mas pode conter todos os bits de um valor de ponteiro. |
| float32 | Número de ponto flutuante IEEE 32 de 754 bits |
| float64 | Número de ponto flutuante IEEE 64 de 754 bits |
| complex64  | Números complexos que contêm flutuação32 como um componente real e imaginário. |
| complex128 | Números complexos que contêm flutuação64 como um componente real e imaginário. |

## Criando novos tipos

Podemos declarar novos tipo de dados com a palavras-chave `type`:

```go
// Alias para tipo string
type Firstname string 

// Alias para tipo time.Time
type Birthdate time.Time

// Alias para tipo map[string]float64 
type ExchangeRate map[string]float64 

// Tipo struct, composto por vário outros tipos
type Hotel struct {   
    Name     string
    Capacity uint8
    Rooms    uint8
    Smoking  bool
}
```

## Conversão de tipos

### String para int

Usando `strconv.ParseInt()`:

```go
func ParseInt(s string, base int, bitSize int) (i int64, err error)
// s: Valor da cadeia de caracteres que deve ser analisado no número inteiro.
// base: O valor base do valor dado, pode ser 0, 2 a 36.
// bitSize: Define o tipo inteiro.
//  0 para int
//  8 para int8
//  16 para int16
//  32 para int32
//  64 para int64

// Exemplo:
strVar := "100"
intVar, err = strconv.ParseInt(strVar, 0, 16)
```

Usando `strconv.Atoi()` (significa ASCII to int). Equivalente a `ParseInt(s, 10, 0)`:

```go
func Atoi(s string) (int, error)
// s: Um valor de cadeia de caracteres que deve ser convertido em um valor inteiro.

// Exemplo:
strVar := "100"
intVar, err := strconv.Atoi(strVar)
```

Usando `fmt.Sscan()` ou `fmt.Sscanf()`:

```go
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
// str: Contém a cadeia de caracteres de argumento a ser extraída.
// format: Formato de cadeia de caracteres com os verbos de formato.
// a: Um tipo personalizado que é usado para especificar um conjunto de uma ou mais assinaturas de método.

// Exemplo:
strVar := "100"
intVar := 0
_, err := fmt.Sscan(strVar, &intVar)

strVar := "id:100"
_, err := fmt.Sscanf(strVar, "id:%3d", &intVar)
```

### String para float

Usando `strconv.ParseFloat()`:

```go
func ParseFloat(s string, bitSize int) (float64, error)
// s: Valor da cadeia de caracteres que deve ser analisado no número flutuante.
// bitSize: Para definir a precisão, o valor pode ser 32 ou 64.

// Exemplo:
s := "-3.141"
floatVar, err := strconv.ParseFloat(s, 8)
```

### String para bool

Usando `strconv.ParseBool()`:

```go
func ParseBool(str string) (bool, error)
//str: String value which is to be parsed in the boolean.

// Exemplo:
s := "true" // ou "t", "T", "1"
boolVar, err := strconv.ParseBool(s)

s := "false" // ou "f", "F", "0"
boolVar, err := strconv.ParseBool(s)
```

### Int para String

Usando `strconv.FormatInt()`:

```go
func FormatInt(i int64, base int) string
// i: Um valor inteiro que deve ser convertido no formato de cadeia de caracteres.
// base: Base do valor dado.

// Exemplo:
var s string = strconv.FormatInt(i, 10)
```

Usando `fmt.Sprintf()`

```go
func Sprintf(format string, a ...interface{}) string
// format: Formato de cadeia de caracteres com os verbos de formato.
// a: Um tipo personalizado que é usado para especificar um conjunto de uma ou mais assinaturas de método.

// Exemplo:
b := 1225
s := fmt.Sprintf("%v", b) // "1225"
```

### Float para String

Usando `strconv.FormatFloat()`:

```go
func FormatFloat(f float64, fmt byte, prec, bitSize int) s string
// f: O número de ponto flutuante que deve ser convertido na forma de cadeia de caracteres.
// fmt: Um valor de byte para definir o formato:
//  'b' (-ddddp±ddd, um expoente binário),
//  'e' (-d.dddde±dd, um expoente decimal),
//  'E' (-d.ddddE±dd, um expoente decimal),
//  'f' (-ddd.dddd, sem expoente),
//  'g' ('e' para grandes expoentes, 'f' caso contrário),
//  'G' ('E' para expoentes grandes, 'f' caso contrário),
//  'x' (-0xd.ddddp±ddd, uma fração hexadecimal e expoente binário), ou
//  'X' (-0Xd.ddddP±ddd, uma fração hexadecimal e expoente binário).
// prec: O valor para definir a precisão que controla o número de dígitos (excluindo o expoente) impressos pelos formatos 'e', 'E', 'f', 'g', 'G', 'x' e 'X'.
//  Para o valor de formato ('e', 'E', 'f', 'x' e 'X') – prec é o número de dígitos após o ponto decimal.
//  E, para o valor de formato ('g' e 'G') – prec é o número máximo de dígitos significativos (zeros à direita são removidos).
//  O valor de precisão especial (-1) é usado para o menor número de dígitos necessários de tal forma que ParseFloat() retorne f exatamente.
// bitSize: Um valor inteiro para definir os bits bitSize (32 para float32, 64 para float64).

// Exemplo:
var f float64 = 3.1415926535
s := strconv.FormatFloat(f, 'E', -1, 32) // "3.1415927E+00"
```

Usando `fmt.Sprintf()`

```go
func Sprintf(format string, a ...interface{}) string
// format: Formato de cadeia de caracteres com os verbos de formato.
// a: Um tipo personalizado que é usado para especificar um conjunto de uma ou mais assinaturas de método.

// Exemplo:
b := 12.454
s := fmt.Sprintf("%v", b) // "12.454"
```

### Bool para String

Usando `strconv.FormatBool()`

```go
func FormatBool(b bool) string
// b: Um valor de bool que deve ser usado para obter "verdadeiro" ou "falso".

// Exemplo:
s := strconv.FormatBool(true)
```

Usando `fmt.Sprintf()`

```go
func Sprintf(format string, a ...interface{}) string
// format: Formato de cadeia de caracteres com os verbos de formato.
// a: Um tipo personalizado que é usado para especificar um conjunto de uma ou mais assinaturas de método.

// Exemplo:
b := true
s := fmt.Sprintf("%v", b) // "true"
```

### Int para Float e Float para Int

Usando `int32()` e `float64()`:

```go
// Exemplo:
var f32 float32 = 10.6556
i32 := int32(f32) // 10
f64 := float64(i32) // 10
```

### Int para Int16, Int32, Int64

Usando `int16()`, `int32()` e `int64()`:

```go
// Exemplo:
var i int = 10
i16 := int16(i)
i32 := int32(i)
i64 := int64(i)
```

### Float32 para Float64 e Float64 para Float32

Usando `float32()` e `float64()`:

```go
// Exemplo:
var f32 float32 = 10.6556
f64 := float64(f32) // 10.6556

f64 = 1097.655698798798
f32 = float32(f64) // 1097.6556
```

## Referências

- <https://www.geeksforgeeks.org/data-types-in-go>
- <https://www.golangprograms.com/go-language/integer-float-string-boolean/how-to-convert-string-to-integer-type-in-go.html>