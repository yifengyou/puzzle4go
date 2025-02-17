package main

import "fmt"

//关于slice或map操作，下面正确的是（）
//A.
//var s []int
//s = append(s,1)
//B.
//
//var m map[string]int
//m["one"] = 1
//C.
//
//var s []int
//s = make([]int, 0)
//s = append(s,1)
//D.
//
//var m map[string]int
//m = make(map[string]int)
//m["one"] = 1
//参考答案：ACD

//B 没有初始化，所以m是nil，赋值，直接抛了,其他的都Ok

func main() {
	var a []int
	a = append(a, 1)
	fmt.Println(a)

	//var b map[string]int
	//b["one"] = 1
	//fmt.Println(b)
	// panic: assignment to entry in nil map
	// map需要先定义，后make申请内存空间

	var c []int
	c = make([]int, 0, 0)
	c = append(c, 1)
	fmt.Println(c)

	var d map[string]int
	d = make(map[string]int)
	d["one"] = 1
	fmt.Println(d)
}
