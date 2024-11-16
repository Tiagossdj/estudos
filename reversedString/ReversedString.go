package main

import "fmt"

// Complete the solution so  that it reverses the string passed into it;
// codeWars kyuu 8

func main() {
	// 'world'  =>  'dlrow'
	// 'word'   =>  'drow'

	fmt.Println(Solution("mata"))

}
func Solution(word string) string {
	var result = ""
	for _, c := range word {
		result = string(c) + result
		fmt.Println(result)
	}
	return result
}
