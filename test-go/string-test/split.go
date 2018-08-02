package main

import (
	"fmt"
	"strings"
)

func getnum(name string) *int {
	a := new(int)
	b := new(int)
	*a = 10
	*b = 20
	if name == "wangyi" {
		return a
	} else {
		return b
	}
}
func main() {
	var m = make(map[string]int)
	clusterNames := strings.Split("wangyi,peng", ",")
	for _, name := range clusterNames {
		num := getnum(name)
		m[name] = *num
	}
	for country := range m {
		fmt.Println(country, m[country])
	}
}
