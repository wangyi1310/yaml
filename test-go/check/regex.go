package main

import "fmt"
import "regexp"

func main() {
	r, _ := regexp.Compile("^[A-Za-z0-9-]+$")
	fmt.Println(r.MatchString(""))

}
