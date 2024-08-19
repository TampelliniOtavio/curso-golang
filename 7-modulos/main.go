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

    pessoa := model.Pessoa {
        Nome: "Nome Legal",
        Endereco: endereco,
    }

    fmt.Println(pessoa)
    fmt.Println(endereco)
    endereco.Numero = 10
    fmt.Println(endereco.Numero)
}
