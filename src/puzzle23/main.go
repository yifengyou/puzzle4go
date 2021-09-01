package main

import "fmt"

func main() {
	//对于局部变量整型切片x的赋值，下面定义正确的是（）
	//A.
	a := []int{
		1, 2, 3,
		4, 5, 6,
	}
	//B.
	//b := []int{
	//	1, 2, 3,
	//	4, 5, 6 // Need a trailing comma before a newline in the composite literal
	//}
	//C.
	//
	c := []int{
		1, 2, 3,
		4, 5, 6}
	//D.
	d := []int{1, 2, 3, 4, 5, 6}

	//x := []int{
	//	1, 2, 3,
	//	4, 5, 6,
	//}

	//x := []int{
	//	1, 2, 3,
	//	4, 5, 6
	//}

	//x := []int{
	//	1, 2, 3,
	//	4, 5, 6}

	x := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(a, c, d, x)
}

//参考答案：ACD
