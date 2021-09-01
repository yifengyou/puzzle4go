package main

import "fmt"

//关于变量的自增和自减操作，下面语句正确的是（）
//A.
//i := 1
//i++
//B.
//
//i := 1
//j = i++
//C.
//
//i := 1
//++i
//D.
//
//i := 1
//i--



// go中只有 x ++ 且属于语句而非表达式
// 左边不能包含任何等号，例如y := x++是不成立的，它是完整的语句而非表达式



func main() {

	//A
	//i := 1
	//i++
	//fmt.Println(i)

	//B j都没有初始化变量，什么鬼赋值给j
	//i := 1
	//j = i++
	//fmt.Println(j)

	//C这是什么鬼操作 ++报错
	//i := 1
	//++i
	//fmt.Println(i)

	//D
	i := 1
	i--
	//j := i++ // ',', ';', new line or '}' expected, got '++'
	fmt.Println(i)
}

//参考答案：AD