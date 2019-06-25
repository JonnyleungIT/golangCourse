package main

import "fmt"

func main() {
	// 声明一个字符串类型变量
	// var username string
	// 标准输入, &:取地址运算符
	// fmt.Scan(&username)
	// fmt.Printf("%s\n",username)

	// 输入多个值,空格/换行作为接收结束
	// var s1,s2 string
	// fmt.Scan(&s1,&s2)
	// fmt.Println(s1,s2)

	// 案例,输入半径值计算圆的半径和周长
	var r float64
	fmt.Scan(&r)
	PI := 3.14159
	fmt.Printf("%.2f\n", 2*PI*r)
	fmt.Printf("%.2f\n", PI*r*r)
}
