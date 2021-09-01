package main

//关于const常量定义，下面正确的使用方式是（）
//A.
//const Pi float64 = 3.14159265358979323846
//const zero = 0.0

//B.
//const (
//	size int64 = 1024
//	eof = -1
//)

//C.
//const (
//	ERR_ELEM_EXIST error = errors.New("element already exists")
//	ERR_ELEM_NT_EXIST error = errors.New("element not exists")
//)

//D.
//const u, v float32 = 0, 3
//const a, b, c = 3, 4, "foo"
//参考答案：ABD

// Go的常量const是属于编译时期的常量，即在编译时期就可以完全确定取值的常量。
// 只支持数字，字符串和布尔，及上述类型的表达式。
// 而切片，数组，正则表达式等等需要在运行时分配空间和执行若干运算才能赋值的变量则不能用作常量。

//const ar = map[string][int]{} // Const initializer 'map[string][int]' is not a constant
//const a = []byte{} // Const initializer '[]byte{}' is not a constant
//const d = func f() {} // Const initializer 'func {...}' is not a constant


//A.
const Pi float64 = 3.14159265358979323846
const zero = 0.0

//B
const (
	size int64 = 1024
	eof        = -1
)

//C.
//const (
//	ERR_ELEM_EXIST error = errors.New("element already exists")
//	// Const initializer 'errors.New("element already exists")' is not a constant
//	ERR_ELEM_NT_EXIST error = errors.New("element not exists")
//	// Const initializer 'errors.New("element already exists")' is not a constant
//)

//D.
const u, v float32 = 0, 3
const a, b, c = 3, 4, "foo"

//const英文本身就是  常量,常数  的意思，那么C ERR_ELEM_EXIST，提示报错  这是什么鬼

func main() {

}
