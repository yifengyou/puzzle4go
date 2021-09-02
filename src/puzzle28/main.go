package main

import "fmt"

//关于GetPodAction定义，下面赋值正确的是（）
type Fragment interface {
	Exec(transInfo *TransInfo) error
}

type TransInfo struct{}
type GetPodAction struct {
}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	//...
	//var fragment Fragment = &GetPodAction{}
	//fmt.Println(fragment)

	//var fragment Fragment = new(GetPodAction)
	//fmt.Println(fragment)

	return nil
}

//A. var fragment Fragment = new(GetPodAction)
//B. var fragment Fragment = GetPodAction
//C. var fragment Fragment = &GetPodAction{}
//D. var fragment Fragment = GetPodAction{}

// 灵魂发问：go接口变量到底传指针还是值
// 接收者类型	接收者值
// 指针	只能为指针
// 值	值与指针类型都可以
// 怎么记？
// 如果接收者是指针，那么如果对应的值非地址，此时就是将值强制转为指针，若值为0就空指针访问，容易出问题。
// 如果接收者是值，本质上指针和值都是数值，只是指针存的是特殊的地址值，值是包含指针的
// 再者，值都是有类型的，那么指针也是一种类型，可以通过反射轻松定位类型，则不管是哪种类型都可以反射转换处理。
// 这里其实就是go编译系统做了工作，智能转换。所有转换都只能一层

//GetPodAction实例，谁可以赋值给接口 Fragment
//A new可以，因为它返回指针类型
//C 很明显了，人家都用&符号了
//B 搞不懂，你把一个结构体赋值过去， 想干啥，什么鬼，完犊子
//D 一个结构体对象赋值给接口变量 ok 啊

func main() {
	i := 0x123456
	fmt.Printf("%T,%#v\n", &i, &i)
	// *int,(*int)(0xc00000a098)
}

//参考答案：ACD
