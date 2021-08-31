package main

import "fmt"

//[intermediate] 关于init函数，下面说法正确的是（）
//A. 一个包中，可以包含多个init函数
//B. 程序编译时，先执行导入包的init函数，再执行本包内的init函数
//C. main包中，不能有init函数
//D. init函数可以被其他函数调用

//C 这个可以有
//D init函数是自动被调用的，不能显示调用

func init() {
	fmt.Println("run init in main")
}

func main() {
	// run init in main
	//init() // Unresolved reference 'init'
}

//参考答案：AB
