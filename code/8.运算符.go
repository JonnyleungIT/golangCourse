package main

import "fmt"

func main() {

	// 算术运算符
	a := 10
	b := 5
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) //整数相除结果为整数
	fmt.Println(a % b)

	// 赋值运算符
	c := 10
	c += 5 // 15
	c -= 5 // 10
	c *= 5 // 50
	c /= 5 // 10
	c %= 5 // 0
	fmt.Println(c)

	// 关系运算符
	person := "李荣浩"

	fmt.Println(person == "周杰伦") // false
	fmt.Println(person != "周杰伦") // true

	var d, e int = 10, 20
	fmt.Println(d < e)  // true
	fmt.Println(d > e)  // false
	fmt.Println(d <= e) // true
	fmt.Println(d >= e) // false

	// 逻辑运算符
	// true && true // true
	// true && false // false
	// true || false // false
	// false || false // false
	// !true // flase

	// 其他运算符
	f := 100
	ptr := &f
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
