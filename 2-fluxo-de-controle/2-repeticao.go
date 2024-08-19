package main

import "fmt"

func main() {
    texto := "palavra"
    tamanho := len(texto)

    for i := 0; i < tamanho; i++ {
        caractere := string(texto[i])

        if caractere == "a" {
            continue
        }

        fmt.Println(caractere)
    }


    // while
    i := 0
    for i < tamanho {
        caractere := string(texto[i])
        i++

        if caractere == "a" {
            continue
        }

        fmt.Println(caractere)
    }
}
