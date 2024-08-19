package model

import (
	"time"
)

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

func (p *Pessoa) IdadeAtual() {
	ano := p.DataDeNascimento.Year()

	anoAtual := time.Now().Year()

    p.Idade = anoAtual - ano
}

func CalculaIdade(p Pessoa) int {
	ano := p.DataDeNascimento.Year()

	anoAtual := time.Now().Year()

	return anoAtual - ano
}
