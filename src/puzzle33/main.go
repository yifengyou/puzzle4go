package main

import (
	"encoding/json"
	"fmt"
)

//golang中大多数数据类型都可以转化为有效的JSON文本，下面几种类型除外（）
//A. 指针
//B. channel
//C. complex
//D. 函数

//JSON 格式是借这个样子滴 {"name":"小红"} ，指针是没问题的
// 链接：https://www.nowcoder.com/questionTerminal/5f062c4dbd004f37841df3733c7ebb01?source=relative
// 来源：牛客网
//
//golang 中的类型比如：channel（通道）、complex（复数类型）、func（函数）均不能进行 JSON 格式化。
//有疑问的地方可能是在A选项指针。
//其实 Pointer（指针）也是能被 JSON 格式化的，
//因为指针会被系统隐式转换为指针所指向的具体对象值，具体的对象值是可以被JSON格式化的。
//BCD怎么变成上面的格式？ 不知道怎么搞了，那就选A吧

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "youyifeng", Age: 28}
	fmt.Printf("%#v\n", p)
	jstr, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Println(string(jstr))
	pp := &Person{Name: "youyifeng", Age: 28}
	fmt.Printf("%#v\n", pp)
	jstr2, err := json.Marshal(pp)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Println(string(jstr2))
}
// main.Person{Name:"youyifeng", Age:28}
//{"Name":"youyifeng","Age":28}
//&main.Person{Name:"youyifeng", Age:28}
//{"Name":"youyifeng","Age":28}



//参考答案：BCD
