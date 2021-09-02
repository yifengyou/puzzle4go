package main

import "fmt"

//9. go语言中的引用类型包含哪些？
// 数组切片、字典(map)、通道（channel）、接口（interface）

func main() {
	arr := []byte{1, 2, 3}
	m := make(map[string]int, 20)
	c := make(chan int, 20)
	fmt.Printf("%T - %T - %T", arr, m, c)

}
