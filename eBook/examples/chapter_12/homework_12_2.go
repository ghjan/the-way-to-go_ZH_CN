package main

/*
练习 12.2: calculator.go

编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
输入的格式为：number1 ENTER number2 ENTER operator ENTER --> 显示结果
当用户输入字符 'q' 时，程序结束。请使用您在练习11.3中开发的 stack 包。
*/

import (
	"bufio"
	"fmt"
	"github.com/ghjan/the-way-to-go_ZH_CN/stack"
	"os"
	"strconv"
	"strings"
)

func main() {
	calculator()
}

func calculator() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")
	mystask := stack.Stack{}
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("There were errors reading, exiting program.")
			return
		}
		token := strings.Replace(input, "\r", "", -1)
		token = strings.Replace(token, "\n", "", -1)
		switch {
		case token == "q": // stop als invoer = "q"
			fmt.Println("Calculator stopped")
			return
		case token >= "0" && token <= "999999":
			i, _ := strconv.Atoi(token)
			mystask.Push(i)
		case token == "+":
			q, _ := mystask.Pop()
			p, _ := mystask.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)+q.(int))
		case token == "-":
			q, _ := mystask.Pop()
			p, _ := mystask.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)-q.(int))
		case token == "*":
			q, _ := mystask.Pop()
			p, _ := mystask.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)*q.(int))
		case token == "/":
			q, _ := mystask.Pop()
			p, _ := mystask.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)/q.(int))
		default:
			fmt.Println("No valid input")

		}

	}
}
