package main

import "fmt"

func main(){

	var date int
	
	fmt.Scanf("%d",&date)
	
	switch date{
	case 1:
		fmt.Println("今天是周一")
	case 2:
		fmt.Println("今天是周二")
	case 3:
		fmt.Println("今天是周三")
	case 4:
		fmt.Println("今天是周四")
	case 5:
		fmt.Println("今天是周五")
	case 6:
		fmt.Println("今天是周六")
	case 7:
		fmt.Println("今天是周日")
	}

	var score float64
	fmt.Print("请输入分数: ")
	fmt.Scan(&score)

	switch int(score) / 10 {
	case 10:
		fmt.Println("完美!")
	case 9:
		fallthrough
	case 8:
		fmt.Println("优秀!")
	case 7:
		fallthrough
	case 6:
		fmt.Println("及格!")
	case 5:
		fallthrough
	case 4:
		fallthrough
	case 3:
		fallthrough
	case 2:
		fallthrough
	case 1:
		fallthrough
	case 0:
		fmt.Println("不及格!")
	default:
		fmt.Println("输入有误!")
	}
}