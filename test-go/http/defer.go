package main

//defer 类似于吧变量做成智能指针一样的东西方便在不用的时候释放掉
// fd socker 还有一些全局的东西

import "fmt"
import "os"
import "io"

func ReadFile() (written int64, err error) {
	fr, err := os.Open("defer.go")
	if err != nil {
		return
	}
	defer fr.Close()

	fw, err := os.Create("defer1.go")
	if err != nil {
		return
	}
	defer fw.Close()
	return io.Copy(fr, fw)
}
func main() {
	ReadFile()
	fmt.Println("copy complete")
}
