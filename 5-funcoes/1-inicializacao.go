package main

import (
	"fmt"
	"strconv"
)

func main() {
    ImprimirMensagem("Hello World")
    ImprimirMensagem("Olá Mundo")
    soma := Soma(1, 3)

    soma2, subtracao, divisao, multiplicacao := Operacao(1, 3)

    ImprimirMensagem(strconv.FormatInt(int64(soma), 10))
    ImprimirMensagem(strconv.FormatInt(int64(soma2), 10))
    ImprimirMensagem(strconv.FormatInt(int64(subtracao), 10))
    ImprimirMensagem(strconv.FormatInt(int64(divisao), 10))
    ImprimirMensagem(strconv.FormatInt(int64(multiplicacao), 10))
}

func ImprimirMensagem(message string) {
    message += ", bom dia!"
    fmt.Println(message)
}

func Soma(numero1 int, numero2 int) int {
    resultado := numero1 + numero2
    return resultado
}

func Operacao(numero1 int, numero2 int) (soma int, subtracao int, divisao int, multiplicacao int) {
    soma = Soma(numero1, numero2)
    subtracao = numero1 - numero2
    divisao = numero1 / numero2
    multiplicacao = numero1 * numero2
    return
}
