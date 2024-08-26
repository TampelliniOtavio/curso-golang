package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
    for i := 0; i < 10000; i++ {
        go showMessage(strconv.Itoa(i)) // go: executa em uma goroutine
    }

    time.Sleep(time.Duration(time.Duration.Seconds(5)))
}

func showMessage(message string) {
    fmt.Println(message)
}
