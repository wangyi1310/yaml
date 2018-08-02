package main

import "fmt"

func function() *int {
	b := 4

	return &b
}
func main() {
	b1 := function()
	fmt.Println(*b1)
}
