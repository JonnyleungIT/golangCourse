package main

import "fmt"

func main() {
	age := 29

	// if语句
	if age < 18 {
		fmt.Println("根据国家有关规定:未成年人限制进入本游戏!")
	}

	// if-else语句
	if age < 18 {
		fmt.Println("根据国家有关规定:未成年人限制进入本游戏!")
	} else {
		fmt.Println("正在进入游戏中...(loading)")
	}

	// if - else if - else
	date := 1
	if date == 7 {
		fmt.Println("周日不限行")
	} else if date%2 == 1 {
		fmt.Println("今天限号: 双号")
	} else {
		fmt.Println("今天限号: 单号")
	}

	// if-else嵌套
	a:=10
	b:=20
	c:=30
	if a > b {
		if b > c {
			fmt.Println("a > b > c")
		}else if a < c{
			fmt.Println("c > a > b")
		}else {
			fmt.Println("a > c > b")
		}
	}else {
		if a > c {
			fmt.Println("b > a > c")
		}else if b < c {
			fmt.Println("c > b > a")
		}else {
			fmt.Println("b > c > a")
		}
	}
}
