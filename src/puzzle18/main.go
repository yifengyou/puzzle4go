package main

import "fmt"

//golang中的指针运算包括（）
//A. 可以对指针进行自增或自减运算
//B. 可以通过“&”取指针的地址
//C. 可以通过“*”取指针指向的数据
//D. 可以对指针进行下标运算
//参考答案：BC

//AD,不知道它在说什么，巴拉巴拉巴拉
// 切片指针，数据指针可以索引，其他指针索引个毛

func main() {
	p := &[]byte{0, 1, 2}
	fmt.Println((*p)[2])
	//p2 := &p
	//fmt.Println(p2[0])
	//Invalid operation: 'p2[0]' (type '**[]byte' does not support indexing)

}
