package main

// No Golang: For e Switch Case
import (
	"fmt"
	"log"
	"os"
)

// global
var (
	cara, coroa int
)

// switch encerra automaticamente após encontrar o primeiro caso correspondente, sem necessidade de um break.
// usar fallthrough se quiser que o proximo caso também seja executado.

func lancarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++
		fmt.Println("Cara:", cara)
	case "coroa":
		coroa++
		fmt.Println("Coroa:", coroa)
	default:
		fmt.Println("caiu em pé!(OMG)")
	}
}

func main() {
	lancarMoeda("lolol")
	a, b := 15, 15

	if a > b {
		fmt.Println("a é maior que b!")
	} else if a < b {
		fmt.Println("a é menor que b!")
	} else {
		fmt.Println("a e b são iguais!")
	}

	// os dois são iguais, mas nesse caso o switch é melhor para visualizar.

	switch {
	case a > b:
		fmt.Println("a é maior que b!")
	case a < b:
		fmt.Println("a é menor que b!")
	default:
		fmt.Println("a e b são iguais!")
	}

	// se eu não tivesse colocado o caminho ESPECÍFICO, ele dava erro!
	file, err := os.Open("Conditionals/hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	// Se não converter para string, o print entrega um array de bytes (data).
	fmt.Println(string(data))

}
