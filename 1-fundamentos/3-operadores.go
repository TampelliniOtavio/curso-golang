package main

import (
	"fmt"
	"reflect"
)

func main() {
    num1 := 10.0
    num2 := 20.0

    text1 := "Texto"
    text2 := "Legal"

    result := num1 / num2
    resultTexto := text1 + text2
    fmt.Println(result)
    fmt.Println(resultTexto)

    fmt.Println(reflect.TypeOf(result))
}
