package main

import (
	"fmt"
	model "golangestudo/model"
	"time"
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
        DataDeNascimento: time.Date(2020, 11, 25, 0, 0, 0, 0, time.Local),
    }

    fmt.Println(pessoa)
    // idade := model.CalculaIdade(pessoa)
    pessoa.IdadeAtual()

    fmt.Println("Idade: ", pessoa.Idade)
}
