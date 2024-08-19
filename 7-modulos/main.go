package main

import (
	"fmt"
	model "golangestudo/model"
)

func main() {
    fmt.Println("iniciando...")

    endereco := model.Endereco {
        Rua: "Rua x",
        Numero: 1,
        Cidade: "São Paulo",
    }

    fmt.Println(endereco)
    endereco.Numero = 10
    fmt.Println(endereco.Numero)
}
