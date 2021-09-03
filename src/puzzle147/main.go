package main

import "fmt"

func main() {
	pase_student()
}

//以下代码有什么问题，说明原因。
type student struct {
	Name string
	Age  int
}

//func pase_student() {
//	m := make(map[string] *student)
//	stus := [] student {
//		{
//			Name: "zhou",
//			Age: 24,
//		}, {
//			Name: "li",
//			Age: 23,
//		}, {
//			Name: "wang",
//			Age: 22,
//		},
//	}
//	for _, stu := range stus {
//		m[stu.Name] = &stu
//	}
//}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{
			Name: "zhou",
			Age:  24,
		}, {
			Name: "li",
			Age:  23,
		}, {
			Name: "wang",
			Age:  22,
		},
	}
	// 错误写法
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
	fmt.Println("-------------------------")
	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}

	// --------------------
	for stu := range stus {
		fmt.Printf("%#v(%p) ", stu, &stu)
	}
	// 0(0xc0000aa058) 1(0xc0000aa058) 2(0xc0000aa058)
	for index, stu := range stus {
		fmt.Printf("\n[%d]%#v(%p)", index, stu, &stu)
	}
	//[0]main.student{Name:"zhou", Age:24}(0xc000096078)
	//[1]main.student{Name:"li", Age:23}(0xc000096078)
	//[2]main.student{Name:"wang", Age:22}(0xc000096078)
}

//考点：foreach
//
//解答：这样的写法初学者经常会遇到的.
//m[stu.Name]=&stu实际上一直指向同一个指针，
//最终该指针的值为遍历的最后一个struct的值拷贝。

// 问题正是出在这里，在for range语句中，v变量用于保存迭代input数组所得的值，但是v只被声明了一次，
// 此后都是将迭代input出的值赋值给v，v变量的内存地址始终未变，这样再将v的地址发送给ch通道，
// 发送的都是同一个地址，当然无法达到预期效果。
