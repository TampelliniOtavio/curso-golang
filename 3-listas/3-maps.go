package main

import "fmt"


func main() {
    // m := map[string]int { "sp": 10000000, "cg": 9000000 }
    m := make(map[string]int)
    m["sp"] = 10000000
    m["cg"] = 9000000

    valor, existe := m["rj"]

    if existe {
        fmt.Println(valor)
    }

    for chave, valor := range m {
        fmt.Println("Cidade: ", chave, " H: ", valor)
    }

    delete(m, "cg")

    fmt.Println(m)
}
