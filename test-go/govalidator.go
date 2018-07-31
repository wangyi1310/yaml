package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

type Server struct {
	ID         string `valid:"uuid,required"`
	Name       string `valid:"machine_id"`
	HostIP     string `valid:"ip"`
	MacAddress string `valid:"mac,required"`
	WebAddress string `valid:"url"`
	AdminEmail string `valid:"email"`
}

func main() {
	server := &Server{
		ID:         "123e4567-e89b-12d3-a456-426655440000",
		Name:       "IX10",
		HostIP:     "127.0.0.1",
		MacAddress: "01:23:45:67:89:ab",
		WebAddress: "www.example.com",
		AdminEmail: "admin@exmaple.com",
	}

	//自定义tag验证函数
	govalidator.TagMap["machine_id"] = govalidator.Validator(func(str string) bool {
		return strings.HasPrefix(str, "IX") //类似于c++中的map只不过第二个间存的是函数指针
	})

	if ok, err := govalidator.ValidateStruct(server); err != nil {
		panic(err)
	} else {
		fmt.Printf("OK: %v\n", ok)
	}
}
