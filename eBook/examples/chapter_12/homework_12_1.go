package main
/*
练习 12.1: word_letter_count.go

编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：
i) 输入的字符的个数，包括空格，但不包括 '\r' 和 '\n'
ii) 输入的单词的个数
iii) 输入的行数
 */
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter （end with S）:")
	input, err := inputReader.ReadString('S')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	count_letters, count_words, count_rows := word_letter_count(strings.Replace(input, "S", "", -1))
	fmt.Printf("count_letters:%d\n", count_letters)
	fmt.Printf("count_words:%d\n", count_words)
	fmt.Printf("count_rows:%d\n", count_rows)

}

func word_letter_count(input string) (int, int, int) {
	rows := strings.Split(input, "\n")
	input = strings.Replace(input, "\r", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	words := strings.Split(input, " ")

	return len(input), len(words), len(rows)
}
