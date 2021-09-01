package main



//[primary] 关于类型转化，下面语法正确的是（）
//A.
//type MyInt int
//var i int = 1
//var j MyInt = i
//B.
//
//type MyInt int
//var i int = 1
//var j MyInt = (MyInt)i
//C.
//
//type MyInt int
//var i int = 1
//var j MyInt = MyInt(i)
//D.
//
//type MyInt int
//var i int = 1
//var j MyInt = i.(MyInt)

//参考答案：C

//A
//type MyInt int
//var i int = 1
//var j MyInt = i //Cannot use 'i' (type int) as the type MyInt

//B.
//type MyInt int
//var i int = 1
//var j MyInt = (MyInt)i //Type 'MyInt' is not an expression

//C
//type MyInt int
//var i int = 1
//var j MyInt = MyInt(i)

//D.
//type MyInt int
//var i int = 1
//var j MyInt = i.(MyInt) // Invalid type assertion: i.(MyInt) (non-interface type int on the left)

func main() {

}
