package main

import (
	"fmt"
)

func main() {
	p := PessoaFisica{
		Pessoa{Nome: "Daniel", Idade: 31, Status: true},
		"Rodrigues",
		"000.000.000.00",
	}
	fmt.Println(p)

}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

type pessoaJuridica struct {
	RazaoSocial string
	cnpj        string
}

func (p PessoaFisica) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos!", p.Pessoa.Nome, p.Pessoa.Idade)
}
