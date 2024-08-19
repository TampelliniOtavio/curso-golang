package main

import (
	"fmt"
	"os"
)

func ShowText() {
    fmt.Println("Finalizando de Manipular Arquivo")
}

func ReadFile() { // Apenas um exemplo, não seria necessario o recover pois seria apenas não executar o panic
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("ReadFile Recuperado")
        }
    }()

    _, err := os.Open("./settings.txt")

    if err != nil {
        panic("FileNotExist")
    }
}

func main() {
    file, err := os.Create("./arquivo-recem-criado.txt")
    defer file.Close() // defer define as últimas instruções de execução
    defer ShowText()

    if err != nil {
        panic(err) // Executar para Casos de Extrema necessidade, pois cancela o processo
    }

    _, err = file.Write([]byte("teste"))

    if err != nil {
        panic(err)
    }

    ReadFile()

    fmt.Println("Fim")
}
