package main

import "fmt"

func main() {
    x := 5
    y := &x

    ImprimirValores(&x, y)
    fmt.Println(x, *y)
    fmt.Println(&x, y)
}

func ImprimirValores(x *int, y *int) {
    *x = 20
}

/*
& -> retorna o ponteiro
* -> busca o valor do ponteiro
*/
