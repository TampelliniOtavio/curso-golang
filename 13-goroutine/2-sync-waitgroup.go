package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(3) // N processos em paralelo
    go callDatabase(&wg)
    go callAPI(&wg)
    go processInternal(&wg)

    wg.Wait()
}

func callDatabase(wg *sync.WaitGroup) {
    time.Sleep(time.Second * 1)
    fmt.Println("Call Database")
    wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
    time.Sleep(time.Second * 2)
    fmt.Println("Call API")
    wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
    time.Sleep(time.Second * 1)
    fmt.Println("Process Internal")
    wg.Done()
}
