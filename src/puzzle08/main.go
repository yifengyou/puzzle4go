package main

import "fmt"

//[primary] 关于循环语句，下面说法正确的有（）
//A. 循环语句既支持for关键字，也支持while和do-while
//B. 关键字for的基本使用方法与C/C++中没有任何差异
//C. for循环支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中断哪一个循环
//D. for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量

// go有关键字洁癖
//A while ,do-while 什么鬼
//B 知道C/C++可以把答案打到屏幕上
// 平行赋值   a,b,c := 1,2,3
// 逗号间隔赋值  a:=1,b:=2,c:=3



/*

强记，go关键字一共就25个
break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var

数据类型、内置函数都是可以重定义，但是一般不这么做
*/

func main() {

	//while x := 100;x<1000 {} // Unresolved reference 'while'

	for i := 0; i < 5; i++ {
		if i == 3 {
			//滚出，循环结束，后面i=4...都不执行
			break
		}
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			//跳出，循环继续执行后面的i=3...继续执行
			continue
		}
	}
	x := 1
	for x < 100 {
		fmt.Println("来啊，快活啊~")
		x++
	}
	for a, b := 1, 2; a < 10; a++ {
		b++
	}

	//for a :=1 , b:="123" ; a < 10; a++ {
	//	b++
	//}

}

//参考答案：CD
