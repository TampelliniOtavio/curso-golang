package main

import "fmt"

func main() {
    slice1 := []int{5, 1, 2, 3}
    slice2 := []string{"a", "e", "f", "g"}
    slice3 := []bool{true, true, false, false, true, false}

    newInts := reverse(slice1)
    newStrings := reverse(slice2)
    newBools := reverse(slice3)

    fmt.Println(slice1, newInts)
    fmt.Println(slice2, newStrings)
    fmt.Println(slice3, newBools)
}

type constraintCustom interface {
    int | string | bool
}

func reverse[T constraintCustom](slice []T) []T {
    newInts := make([]T, len(slice))

    newIntsLen := len(slice) - 1
    for i := 0; i < len(slice); i++ {
        newInts[newIntsLen] = slice[i]
        newIntsLen--
    }

    return newInts
}
