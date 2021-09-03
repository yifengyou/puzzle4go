package main

import (
	"fmt"
)

// map多种初始化方式
// 赋值初始化 m := map[int32]int32{}
// make m := make(map[int32]int32)
// new m := new(map[int32]int32)

func main() {
	m1 := map[int32]int32{}
	fmt.Printf("%#v-%p-%T\n", m1, m1, m1)
	m1[1] = 2
	fmt.Printf("%#v-%p-%T\n", m1, m1, m1)
	// ---------------------------------
	m2 := make(map[int32]int32)
	fmt.Printf("%#v-%p-%T\n", m2, m2, m2)
	m2[1] = 2
	fmt.Printf("%#v-%p-%T\n", m2, m2, m2)
	// ---------------------------------
	m3 := new(map[int32]int32)
	fmt.Printf("%#v-%p-%T\n", m3, m3, m3)
	//m3[1] = 2
	//// invalid operation: m3[1] (type *map[int32]int32 does not support indexing)
	*m3 = make(map[int32]int32)
	(*m3)[1] = 1
	fmt.Printf("%#v-%p-%T\n", (*m3)[1], (*m3)[1], (*m3)[1])
	fmt.Printf("%#v-%p-%T\n", m3, m3, m3)

	var m4 map[int32]int32
	m4 = map[int32]int32{}
	fmt.Printf("%#v-%p-%T\n", m4, m4, m4)

	m5 := new(map[int32]int32)
	*m5 = make(map[int32]int32)
	(*m5)[1] = 1
	fmt.Printf("%#v - %T\n", m5, m5)

	m6 := new(int)
	*m6 = 1
	fmt.Printf("%#v - %T\n",*m6, *m6)
	fmt.Printf("%#v - %T\n",m6, m6)

	// panic: assignment to entry in nil map
}

//map[int32]int32{}-0xc00007a480-map[int32]int32
//map[int32]int32{1:2}-0xc00007a480-map[int32]int32
//map[int32]int32{}-0xc00007a510-map[int32]int32
//map[int32]int32{1:2}-0xc00007a510-map[int32]int32
//&map[int32]int32(nil)-0xc000006030-*map[int32]int32
//1-%!p(int32=1)-int32
//&map[int32]int32{1:1}-0xc000006030-*map[int32]int32
//map[int32]int32{}-0xc00007a600-map[int32]int32
//&map[int32]int32{1:1} - *map[int32]int32
//1 - int
//(*int)(0xc00000a0e8) - *int
//
//Process finished with the exit code 0

// 考点 ： map初始化

// 在go语言中，可以使用new和make来创建map类型的变量，但是它们之间有很明显的不同，
// 使用new来创建map时，返回的内容是一个指针，这个指针指向了一个所有字段全为0的值map对象，
// 需要初始化后才能使用，而使用make来创建map时，返回的内容是一个引用，可以直接使用。

// https://blog.csdn.net/choumin/article/details/89893830
