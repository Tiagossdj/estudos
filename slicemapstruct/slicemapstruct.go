package main

import (
	"fmt"
)

// slice, map e struct
// len para saber tamanho do slice e cap para sua capacidade

/*o Go aumenta a capacidade do slice na hora do append, nesse caso, de 4 para 8, ele ja prepara para possiveis adicões*/

func main() {
	/*nomes := []string{"Matias", "Nathan", "Maria", "Nicole"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Felipe")
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Lucas")
	fmt.Println(nomes, len(nomes), cap(nomes))

	// Com o append de lucas a capacidade ainda é 8, e ira continuar até chegar ao limite.
	*/

	/*
		nomes := make([]string, 10, 20)
		fmt.Println(nomes)

		idades := make(map[string]uint8)
		idades["Daniel"] = 32
		idades["Aline"] = 34
		idades["Isabel"] = 42

		val, ok := idades["Daniel"]
		fmt.Println(val, ok)
	*/

	type Pessoa struct {
		Nome      string
		Sobrenome string
		Idade     int8
		Status    bool
	}

	p := Pessoa{
		Nome:      "Tales",
		Sobrenome: "Rodrigues",
		Idade:     30,
		Status:    true,
	}

	fmt.Println(p)
}
