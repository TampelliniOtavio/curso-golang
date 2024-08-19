package main

import (
	"errors"
	"fmt"
	"math"
)

type Geometria interface {
	area() float64
}

type Retangulo struct {
	Largura, Altura float64
}

func (ret Retangulo) area() float64 {
	return ret.Altura * ret.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) area() float64 {
	return c.Raio * c.Raio * math.Pi
}

func ExibirArea(g Geometria) {
	fmt.Println(g.area())
}

type ProblemaDeNetwork struct {
	rede     bool
	hardware bool
}

func (p ProblemaDeNetwork) Error() string {
	if p.rede {
		return errors.New("Problema de Rede").Error()
	}

	if p.hardware {
		return errors.New("Problema de hardware").Error()
	}

	return "Outro Problema"
}

func ExibirError(err error) {
	fmt.Println(err.Error())
}

func main() {
	fmt.Println("inicializando...")

	retangulo := Retangulo{
		Largura: 1,
		Altura:  2,
	}

	cirulo := Circulo{
		Raio: 3,
	}

	ExibirArea(retangulo)

	ExibirArea(cirulo)

	p := ProblemaDeNetwork{
		rede:     true,
		hardware: false,
	}

	ExibirError(p)

	var lista []interface{}
	lista = append(lista, 10)
	lista = append(lista, 8.123)
	lista = append(lista, "texto")
	lista = append(lista, true)

	for _, valor := range lista {
		if v, ok := valor.(string); ok { // instancia a variável e utiliza no if
			fmt.Println(v + " string")
		} else {
			fmt.Println(valor)
		}
	}

}
