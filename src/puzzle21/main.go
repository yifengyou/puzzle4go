package main

import "fmt"

//关于整型切片的初始化，下面正确的是（）
//A. s := make([]int)
//B. s := make([]int, 0)
//C. s := make([]int, 5, 10)
//D. s := []int{1, 2, 3, 4, 5}

// A 没有长度
// 务必注意，创建无缓冲channel可以没有size，例如make(chan int)
// 内建函数 make **只能**用来为 slice， map 或 chan 类型分配内存和初始化一个对象

func main() {
	//s := make([]int) // Missing len argument in the make function
	// func make(t Type, size ...IntegerType) Type
	a := make([]int, 0)
	b := make([]int, 5, 10)
	c := []int{1, 2, 3, 4, 5}
	ch := make(chan int)

	fmt.Println(a, b, c, cap(ch))
}

//参考答案：BCD
