package main

import "fmt"

func main() {
    lista := []int {2, 8, 3, 10, 5, 4, 7, 9, 1}

    var somaAte5 int
    var somaAcimaDe5 int

    for i := 0; i < len(lista); i++ {
        campo := lista[i]

        if campo <= 5 {
            somaAte5 += campo
            continue
        }

        somaAcimaDe5 += campo
    }

    fmt.Println("Soma entre 1 e 5: ", somaAte5)
    fmt.Println("Soma entre 6 e 10: ", somaAcimaDe5)
}

/*
Dado um slice com os itens "2, 8, 3, 10, 5, 4, 7, 9, 1" que vão de 1 a 10,
efetuar a soma em duas variáveis, a primeira números de 1 a 5 e a segunda de 6 a 10.
Imprimir os dois resultados
*/
