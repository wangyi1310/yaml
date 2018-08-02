package main

import "fmt"
import "strconv"

//主要是提供字符串转换函数
//https://studygolang.com/articles/5003 所有能转换的东西都在这

func main() {
	i, err := strconv.ParseInt("12sdfasd", 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
