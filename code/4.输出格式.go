package main

import "fmt"

func main() {

	// fmt.Println():标准输出,打印时自动换行
	fmt.Println("hello")
	fmt.Println(3.14159)
	fmt.Println(188)

	// fmt.Print():打印时不会换行
	fmt.Print("jonny\n")
	fmt.Print("leung\n")

	// fmt.Printf():按照格式输出,输出时不会换行
	a := 10
	b := 3.14
	fmt.Printf("%-5d\n", a) // 左对齐,宽度5
	fmt.Printf("%07d\n", a) // 右对齐,宽度7,空位补0
	fmt.Printf("%.2f\n", b) // 保留2位小数

}
