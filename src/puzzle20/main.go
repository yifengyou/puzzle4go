package main

import "fmt"

//下面赋值正确的是（）
//A. var x = nil
//B. var x interface{} = nil
//C. var x string = nil
//D. var x error = nil

// 非引用类型无法直接nil赋值
// 指针，引用类型的变量的默认值是nil

func main() {
	//var x = nil // Cannot assign nil without the explicit type

	//var y = []byte(nil)
	//fmt.Println(y)

	//var i interface{} = nil
	//var s string = nil  // Cannot use 'nil' as the type string
	var x error = nil

	//A 不能用nil给一个没有类型的变量
	//C string不能nil赋值
	fmt.Println(x)

	// error是一个接口
	//type error interface {
	//	Error() string
	//}

}

//参考答案：BD
