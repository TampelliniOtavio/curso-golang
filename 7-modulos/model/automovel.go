package model

type Automovel struct {
	Ano    int
	Placa  string
	Modelo string
}

type Moto struct {
	Automovel
	Cilindradas int
}

type Carro struct {
	Automovel
	QuantidadeDePortas   int
	Potencia             int
	PossuiArCondicionado bool
}
