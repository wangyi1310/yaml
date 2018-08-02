package main

import "fmt"

type test struct {
	Id   string
	Name string
}

func (t *test) construct() {

	fmt.Println("constructting")
}

func (t *test) destruct() {
	fmt.Println("destructting")
}
func main() {

	var t test
	t.construct()
	t.destruct()

	t1 := new(test)
	t1.construct()
	t1.destruct()
}
