package math


func Soma(valores ...int) (total int) {
    for _, valor := range valores {
        total += valor
    }
    return
}

func Subtracao(valores ...int) (total int) {
    for _, valor := range valores {
        total -= valor
    }
    return
}