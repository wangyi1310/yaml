package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var n string
	fmt.Println("please input sting")
	fmt.Scanf("%s", &n)
	var cmd string
	//cmd = fmt.Sprintf("select * from jiqi where name like '%%%s%%'", n)
	cmd = fmt.Sprintf("select * from jiqi where name like %s", n)
	fmt.Printf("%s", cmd)
	rows, _ := db.Query(cmd)
	//rows, err := db.Query("select * from jiqi where name like '%%%?%%'", n)
	name := ""
	for rows.Next() {
		erri := rows.Scan(&name)
		if erri != nil {
			fmt.Println("dsafdsf")
		}
		fmt.Println(name)
	}
	//fmt.Println(name)
	//var b int = 10
	//var c int = 12
	//var m = make(map[string]int)
	//m["1"] = b
	//m["2"] = c
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

}
