package main

import "fmt"

//[primary] 下面属于关键字的是（）
//A. func
//B. def  python的   干扰我们的
//C. struct
//D. class java的	干扰我们的


func main() {
	int := 123
	fmt.Println(int) // 123
	float64 := "abc"
	fmt.Println(float64) // abc
}

/*

强记，go关键字一共就25个
break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var

数据类型、内置函数都是可以重定义，但是一般不这么做
*/



//参考答案：AC