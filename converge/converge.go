package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := converge(trabalho("Sol"), trabalho("Lua"))

	for x := 0; x < 15; x++ {
		fmt.Println(<-canal)
	}
}

func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Função %v diz:%v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)
	return canal
}

func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()

	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}

// 	par := make(chan int)
// 	impar := make(chan int)
// 	converge := make(chan int)

// 	go envia(par, impar)
// 	go recebe(par, impar, converge)

// 	for v := range converge {
// 		fmt.Println("Valor recebido:", v)
// 	}
// }

// func envia(p, i chan int) {
// 	x := 15
// 	for n := 0; n < x; n++ {
// 		if n%2 == 0 {
// 			p <- n
// 		} else {
// 			i <- n
// 		}
// 	}
// 	close(p)
// 	close(i)
// }

// func recebe(p, i, c chan int) {

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {

// 		for v := range p {
// 			c <- v
// 		}
// 		wg.Done()
// 	}()
// 	wg.Add(1)
// 	go func() {

// 		for v := range i {
// 			c <- v
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	close(c)
