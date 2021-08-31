package main

import "fmt"

//[primary] 定义一个包内全局字符串变量，下面语法正确的是 （）
//A. var str string
//B. str := ""
//C. str = ""
//D. var str = ""
//参考答案 AD


/*
1. 短变量只能是局部变量
2. 全局变量定义时赋值可以不用指定类型，会自动识别类型，若没有定义时赋值则需要类型
 */
var str string
//var str = ""
//str := ""
//str = ""
func main() {
	fmt.Println(str)
}




//参考答案：AB