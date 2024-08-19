package main

import "fmt"

func main() {
    salario := 500.00
    var salarioMaisBonus float32
    tipo := "CLT"
    
    salarioMaisBonus = float32(salario)

    if salario > 1000 {
        salarioMaisBonus += 100
    } else if salario <= 500 && tipo == "CLT" {
        salarioMaisBonus += 500
    } else {
        salarioMaisBonus += 300
    }

    fmt.Println("Salário: ", salarioMaisBonus)
}
