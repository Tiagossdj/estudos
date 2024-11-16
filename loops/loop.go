package main

import (
	"fmt"
)

// Loop em Go só o FOR, não tem while!

func main() {
	/*var lerNumero int
	fmt.Print("Digite o número de loops: ")
	fmt.Scan(&lerNumero)
	for i := 0; i < lerNumero; i++ {
		fmt.Println(i)*/

	nomes := []string{"ricardo", "Daniel", "mariana", "katia"}
	var i int
	for i < len(nomes) {
		fmt.Println(nomes[i])
		i++
	}
}
