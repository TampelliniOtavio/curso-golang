package model

import (
	"errors"
	"time"
)

type ComprasDoMes struct {
	DataDaCompra     time.Time
	Mercado          string
	ItensParaComprar []Produto
}

type Produto struct {
	Nome string
}

func NovaCompra(mercado string, itens []string) (*ComprasDoMes, error) {
    if mercado == "" {
        return nil, errors.New("Mercado é Obrigatório")
    }

    if len(itens) == 0 {
        return nil, errors.New("Pelo menos um produto no carrinho")
    }

    produtos := make([]Produto, 0)
    
    for i := 0; i < len(itens); i++ {
        produtos = append(produtos, Produto{
            Nome: itens[i],
        })
    }

	return &ComprasDoMes{
		DataDaCompra:     time.Now(),
		Mercado:          mercado,
		ItensParaComprar: produtos,
	}, nil
}
