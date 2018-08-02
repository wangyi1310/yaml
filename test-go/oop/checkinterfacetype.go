package main

import "fmt"

//interface 接口中定义的变量都是 public final and static
func main() {
	var a interface{} = "helloworld"
	var b = 10
	a = b
	_, ok := a.(int)
	if ok {
		fmt.Println("该变量是int")
	}
}
