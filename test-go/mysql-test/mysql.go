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
	//cmd := fmt.Sprintf("select * from jiqi where name like '%%%s%%'", n)

	//rows, err := db.Query(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	name := ""
	for rows.Next() {
		rows.Scan(&name)
		fmt.Println(name)
	}
}
