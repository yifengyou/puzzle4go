package main

import (
	"encoding/json"
	"fmt"
)

//关于map，下面说法正确的是（）
//A. map反序列化时json.unmarshal的入参必须为map的地址
//B. 在函数调用中传递map，则子函数中对map元素的增加不会导致父函数中map的修改
//C. 在函数调用中传递map，则子函数中对map元素的修改不会导致父函数中map的修改
//D. 不能使用内置函数delete删除map的元素

func main() {
	m := make(map[string]int, 100)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	mj, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal error", err)
		return
	}
	fmt.Println(string(mj))
	//buff := make([]byte, 1000)
	buff := make(map[string]int)
	err = json.Unmarshal(mj, &buff)
	if err != nil {
		fmt.Println("unmashal error", err)
		return
	}
	fmt.Println(buff)
	fmt.Println(buff["a"])
	fmt.Println(buff["b"])
	fmt.Println(buff["c"])
}

//参考答案：A
