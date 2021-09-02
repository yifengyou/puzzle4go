package main

//flag是bool型变量，下面if表达式符合编码规范的是（）
//A. if flag == 1
//B. if flag
//C. if flag == false
//D. if !flag


// go语言是强类型


var flag bool
func main() {
	//if flag == 1 // Invalid operation: flag == 1 (mismatched types bool and untyped int)
	if flag{}
	if flag == false{}
	if !flag{}
}


//参考答案：BCD