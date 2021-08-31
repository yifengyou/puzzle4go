package main

import "fmt"

//[primary] 关于字符串连接，下面语法正确的是（）
//A. str := ‘abc’ + ‘123’
//B. str := "abc" + "123"
//C. str ：= '123' + "abc"
//D. fmt.Sprintf("abc%d", 123)

func main() {
	//str := ‘abc’ + ‘123’ // 123 evaluated but not used
	str := "abc" + "123"
	//str := '123' + "abc" // Too many characters in the rune literal
	s := fmt.Sprintf("abc%d", 123)
	// String printf return string type

	fmt.Println(str) // abc123
	fmt.Println(s) // abc123
}

//参考答案：BD
