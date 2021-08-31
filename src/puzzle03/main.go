package main

import "fmt"

//[primary] 通过指针变量 p 访问其成员变量 name，下面语法正确的是（）
//A. p.name
//B. (*p).name   *p.name
//C. (&p).name   &p.name
//D. p->name
//

type Person struct {
	name string
}

func main() {
	p := &Person{name: "小红"}
	fmt.Println(p.name)
	fmt.Println((*p).name)
	//fmt.Println(*p.name)  // Invalid indirect of 'p.name' (type 'string')
	fmt.Println(&p)  // 0xc000006028
	fmt.Println(&p.name)  // 0xc000060240
	//
	//对指针再次取地址，什么鬼
	//fmt.Println(&p.name)
	//c语言的写法
	//fmt.Println(p->name)
}



//参考答案：AB