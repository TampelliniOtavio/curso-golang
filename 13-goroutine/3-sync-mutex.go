package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    var m sync.Mutex
    i := 0

    for x := 0; x < 10000; x++ {
        go func() {
            m.Lock()
            defer m.Unlock()
            i++
        }()
    }

    time.Sleep(time.Second * 1)
    fmt.Println(i)
}
