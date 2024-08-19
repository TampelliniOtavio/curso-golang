package main

import "fmt"

type endereco struct {
    rua string
    numero int
    cidade string
}

func main() {
    fmt.Println("iniciando...")

    endereco := endereco {
        rua: "Rua x",
        numero: 1,
        cidade: "São Paulo",
    }

    fmt.Println(endereco)
    endereco.numero = 10
    fmt.Println(endereco.numero)
}
