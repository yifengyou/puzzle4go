package main
import "C"

//value是整型变量，下面if表达式符合编码规范的是（）
//A. if value == 0
//B. if value
//C. if value != 0
//D. if !value


// go语言是强类型，condition必须为bool类型
// BD都是布尔类型直接判断的

var value int
func main() {
	if value == 0{}
	//if value{} // The non-bool value 'value' (type int) used as a condition
	if value != 0{}
	//if !value{} // Invalid operation: ! int
}


//参考答案：AC