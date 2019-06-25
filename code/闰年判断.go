// 判断输入的年份是否为闰年

package main

import "fmt"

func main() {
	var year int
	fmt.Scanf("%d",&year)
	if (year%400) == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("您输入的年份为闰年.")
	} else {
		fmt.Println("您输入的年不是闰年.")
	}
}
