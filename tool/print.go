package tool

import "fmt"

/*
	终端彩色输出：https://zhuanlan.zhihu.com/p/76751728
*/

func PrintGreen(str string) {
	fmt.Printf("\033[0;32;40m%s\033[0m\n", str)
}

func PrintYello(str string) {
	fmt.Printf("\033[0;33;40m%s\033[0m\n", str)
}

func PrintRed(str string) {
	fmt.Printf("\033[1;31;40m%s\033[0m\n", str)
}
