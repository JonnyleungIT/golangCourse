package main

import "fmt"

func main() {
	const GameName string = "植物大战僵尸之逆袭"
	fmt.Println(GameName)

	var a int = 88
	const b int = 188
	fmt.Println(a + b) // 常量可以与变量进行计算

	// iota枚举
	const (
		c = iota + 10000
		d
		e
		f
		g
		h
	)
	fmt.Println(c, d, e, f, g)

	// iota写在同一行值相同
	const (
		i, j = iota, iota
	)
	fmt.Println(i, j)

	// iota遇const置0

}
