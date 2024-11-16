package main

import (
	"fmt"
)

type Categoria struct {
	Nome string
	Pai  *Categoria
}

func main() {
	cat := Categoria{Nome: "Categoria 1"}
	cat.nomePai(&Categoria{Nome: "Pai"})
	if !cat.HasParent() {
		fmt.Println("não tem pai")
	} else {
		fmt.Println("Tem Pai!")
	}

	fmt.Println(cat)

}

func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

func (c *Categoria) nomePai(pai *Categoria) {
	c.Pai = pai

}
