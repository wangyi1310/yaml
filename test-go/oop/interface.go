package main

import (
	"fmt"
)

type Phone interface {
	call() string
}
type N struct {
}
type I struct {
}

func (i I) call() string {
	fmt.Println("I")
	return string("helloworld")
}
func (n N) call() string {
	fmt.Println("N")
	return string("helloworld")
}
func main() {
	var phone Phone //类似于c++ 中的虚函数一样的东西
	phone = new(N)
	fmt.Println(phone.call())
}
