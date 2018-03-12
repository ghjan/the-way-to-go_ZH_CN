package main

import (
	"fmt"
	"os"
)

func main() {
	//输出到屏幕 WriteString是File的方法
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0)
	defer f.Close()
	f.WriteString("hello, world in a file\n")
	fmt.Fprintf(f, "some test data.\n")

}
