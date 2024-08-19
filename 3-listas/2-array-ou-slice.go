package main

import "fmt"

func main() {
    var listaToda = []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
    listaToda = append(listaToda, 10)

    fmt.Println(listaToda)
    fmt.Printf("%T\n", listaToda)

    var listaArray = [11]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
    // Ocorre um erro, pois o Array não é mutável
    // listaArray = append(listaArray, 10)

    fmt.Println(listaArray)
    fmt.Printf("%T\n", listaArray)
}
