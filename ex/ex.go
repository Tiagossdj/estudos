package main

func Main() {
}

func Soma(t ...int) int {
	total := 0
	for _, v := range t {
		total += v
	}
	return total
}
