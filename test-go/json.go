package main

import (
	"encoding/json"
	"fmt"
)

// 所有json格式的默认变量第一个字母都要大写
type Product struct {
	Name   string
	Id     int64
	Price  float64
	Online bool
}

//或者打key不用变量名作为标签的key值

type Product1 struct {
	Name   string  `json:"name"`
	Id     int64   `json:"id"`
	Price  float64 `json:"price"`
	Online bool    `json:"online"` // 如果后面是`json: "-"` 表示该标签不导出omitempy表示不导出为0的标签
}

func main() {
	p := &Product1{}
	p.Name = "red mi 5"
	p.Online = true
	p.Price = 50.3
	p.Id = 13108
	data, erro := json.Marshal(p)
	if erro != nil {
		fmt.Println("json marshal error")
	}
	fmt.Println(string(data))
	p1 := &Product1{}
	if err := json.Unmarshal(data, p1); err != nil {
		fmt.Println("json unmarshal erorr")
		fmt.Println(err)
	}
	fmt.Println(p1.Name)
}
