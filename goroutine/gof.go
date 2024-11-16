package main

import (
	"fmt"
	"time"
)

func numeros(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Escrito no channel: %d\n", i)
		time.Sleep(time.Millisecond * 90)
	}
	fmt.Println("Fim da Escrita")
	close(n)
}

/* func letras(done chan<- bool) {
for l := 'a'; l < 'j'; l++ {
	fmt.Printf("%c ", l)
	time.Sleep(time.Millisecond * 230)
}
done <- true
}
*/

// / channels (sem buffer) para remover o sleep
// / channel com buffer
func main() {

	cn := make(chan int, 3)
	go numeros(cn)

	for v := range cn {
		fmt.Printf("lido do channel: %d\n", v)
		time.Sleep(time.Millisecond * 180)

	}

	fmt.Println("Fim da Execução!")
}
