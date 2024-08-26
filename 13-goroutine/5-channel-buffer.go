package main

import (
	"fmt"
	"time"
)

func main() {
    channel := make(chan int, 100) // Com buffer de 100
    go setListChannel(channel)

    for v := range channel {
        fmt.Println("Recebendo: ", v)
        time.Sleep(time.Second)
    }
}
 // chan<-: apenas escrita
 // <-cha: apenas leitura
func setListChannel(channel chan<- int) {
    for i := 0; i < 100; i++ {
        channel <- i
        fmt.Println("Enviando: ", i)
    }

    close(channel)
}
