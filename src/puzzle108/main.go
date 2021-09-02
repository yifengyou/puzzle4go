package main

import "fmt"

//说说go语言中的for循环？
// 1.for循环支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中断哪一个循环
// 2 for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量
// 3 最简单的for循环 for{}  死循环
// 4 for常与range组合,且有很骚气的下表展开问题

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	m := make(map[string]int, 20)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	for key, value := range m {
		fmt.Println(key, value)
	}
	for key := range m {
		fmt.Println(key)
	}

	for {
	}
}
