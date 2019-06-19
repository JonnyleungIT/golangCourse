package main

import "fmt"

func main(){
	
	// 布尔
	var bool1 bool
	fmt.Println(bool1)

	// 浮点
	var f1 float32 = 3.1415926535897932384626
	var f2 float64 = 3.1415926535897932384626
	f3 := 3.14 // 浮点自动推导为float64类型
	fmt.Println(f1,f2,f3)

	// 字符
	var a byte = 'a'
	fmt.Println(a) // 默认打印字符对应ascii码值

	// 字符串
	var str string
	var slogan1 string = "我爱学习"
	var slogan2 string = "学习使我快乐"
	fmt.Println(str) // 字符串零值:\0
	fmt.Println(slogan1+slogan2) // 字符串相加操作
	
	// 
}