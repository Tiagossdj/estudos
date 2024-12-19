package main

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{10, 11, 22}, 43},
		{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		z := Soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "got:", z)
		}
	}
}

func ExampleSoma() {
	fmt.Println(Soma(3, 2, 1))
	// Output: 6
}

func TestSoma(t *testing.T) {
	teste := Soma(1, 3)
	res := 4
	if teste != res {
		t.Error("Expected:", res, "got:", teste)
	}
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(1, 1)
	}

}

func BenchmarkSomaEmTabela(b *testing.B) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{10, 11, 22}, 43},
		{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		Soma(v.data...)
	}
}
