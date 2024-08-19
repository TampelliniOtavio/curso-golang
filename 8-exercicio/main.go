package main

import (
	model "exercicio/model"
	"fmt"
)


func main() {
    compras, error := model.NovaCompra("Sonda", []string{"Banana", "Maçã"})

    if error != nil {
        fmt.Println(error.Error())
    } else {
        fmt.Println(*compras)
    }
}
