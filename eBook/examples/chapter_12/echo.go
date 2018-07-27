package main

import (
	"flag" // command line option parser
	"os"
)
//定义了一个默认值是 false 的 flag
var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	//打印 flag 的使用帮助信息
	flag.PrintDefaults()
	/*
	flag.Parse() 扫描参数列表（或者常量列表）并设置 flag, flag.Arg(i) 表示第i个参数。
	Parse() 之后 flag.Arg(i) 全部可用，flag.Arg(0) 就是第一个真实的 flag，而不是像 os.Args(0) 放置程序的名字。
	 */
	flag.Parse() // Scans the arg list and sets up flags
	var s = ""
	for i := 0; i < flag.NArg(); i++ { //flag.Narg() 返回参数的数量
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
