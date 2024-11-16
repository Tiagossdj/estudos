package main

import (
	"fmt"
)

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}
type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu cpf é: %s", pf.cpf)
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos!", p.Nome, p.Idade)
}

type pessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string
}

func (pj pessoaJuridica) Doc() string {
	return fmt.Sprintf("meu cnpj é: %s", pj.cnpj)
}

type Document interface {
	Doc() string
}

func show(d Document) {
	switch d.(type) {
	case PessoaFisica:
		fmt.Println(d.(PessoaFisica).Sobrenome)
	case pessoaJuridica:
		fmt.Println(d.(pessoaJuridica).RazaoSocial)
	default:
		fmt.Println("Tipo não Identificado")
	}

	fmt.Println((d))
	fmt.Println((d.Doc()))
}

func main() {
	pf := PessoaFisica{
		Pessoa{Nome: "Daniel", Idade: 31, Status: true},
		"Rodrigues",
		"000.000.000.00",
	}
	show(pf)
	println()

	pj := pessoaJuridica{
		Pessoa{Nome: "Learn Go", Idade: 0, Status: true},
		"TT LTDA",
		"00.000.000/0000-00",
	}
	show(pj)

}
