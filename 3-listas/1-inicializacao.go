package main

import "fmt"

func main() {
    lista := []int{1,4,7,2,5,7}
    lista = append(lista, 45)

    listaDeString := []string {"a", "b", "c"}
    listaDeString = append(listaDeString, "d")

    listaSemValores := make([]int, 0)

    fmt.Println(lista)
    fmt.Println("Indice 1: ", lista[0])
    fmt.Println("Tamanho: ", len(lista))

    fmt.Println(listaDeString)
    fmt.Println("Indice 1: ", listaDeString[0])
    fmt.Println("Tamanho: ", len(listaDeString))

    fmt.Println("Lista Inicializada: ", listaSemValores)

    listaCortada := lista[:3] // 3 primeiros valores

    listaCortada2 := lista[3:] // a partir do indice 3
    ultimoItem := lista[len(lista)-1:]

    fmt.Println("Lista Cortada: ", listaCortada)
    fmt.Println("Lista Cortada 2: ", listaCortada2)
    fmt.Println("Lista Cortada Ultimo Item: ", ultimoItem)
}
