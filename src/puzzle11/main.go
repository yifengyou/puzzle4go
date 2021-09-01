package main

import "fmt"

//关于局部变量的初始化，下面正确的使用方式是（）
//A. var i int = 10
//B. var i = 10
//C. i := 10
//D. i = 10


// 短变量赋值仅用于局部变量
// 所有全局变量赋值方式都可以作为局部变量的赋值，反之不成立

func main() {
	var i int = 10
	//var i = 10
	//i := 10
	//i = 10

	fmt.Println(i)
}

//参考答案：ABC
