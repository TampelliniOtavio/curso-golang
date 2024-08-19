package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
    var numero int8 = 127
    var numeroInt int
    numeroInt = int(numero)

    var numeroFloat float32
    numeroFloat = float32(numero)

    fmt.Println(numero, numeroInt, numeroFloat)
    fmt.Println(reflect.TypeOf(numero), reflect.TypeOf(numeroInt), reflect.TypeOf(numeroFloat))

    // DOCS: https://pkg.go.dev/strconv

    _, erra := strconv.ParseBool("a")
    b, errb := strconv.ParseBool("true")

    fmt.Printf("%s \n", erra)
    fmt.Println(b, errb)
}

