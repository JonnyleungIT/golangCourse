package main

import "fmt"

func main() {

	// 变量的声明
	var name string
	var age int
	var gender string

	// 变量赋值
	name = "jonnyleung"
	age = 22
	gender = "男"
	fmt.Println(name, age, gender)

	// 变量的定义
	var id int = 2014122133
	fmt.Println(id)

	// 自动推导类型
	location := "广州" // string
	year := 1995     // int
	price := 5.6     // float64
	char := 'A'      // byte
	status := true   // bool
	fmt.Println(location, year, price, char, status)

	// 案例1: 计算圆的周长
	var radius float64
	var long float64
	var PI float64 = 3.14159
	radius = 2.4
	long = 2 * PI * radius
	fmt.Println(long)
}
