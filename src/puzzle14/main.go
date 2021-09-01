package main

import "fmt"

//下面的程序的运行结果是（）
//func main() {
//	if (true) {
//		defer fmt.Printf("1")
//	} else {
//		defer fmt.Printf("2")
//	}
//	fmt.Printf("3")
//}
//A. 321
//B. 32
//C. 31
//D. 13


//A 人家有一个if，怎么可能是3个
//defer是在函数执行后，再执行，所以先执行3，然后是true，那么就是1
//defer需要注册才能使用，执行到defer语句的时候其实就是注册过程
//defer是有性能损耗的

func main() {
		if (true) {
			defer fmt.Printf("1")
		} else {
			defer fmt.Printf("2")
		}
		fmt.Printf("3")
		// 31
}


//参考答案：C