package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

//69 数组是一个值类型（）
//参考答案：对的

//70 使用map不需要引入任何库（）
//参考答案：对的

//71 内置函数delete可以删除数组切片内的元素（）
//参考答案：错的  delete仅用于删除map的键值对

//72 指针是基础类型（）
//参考答案：错的
//
//整数类型： int8、int16等等
//浮点类型：float32、float64
//布尔类型：bool
//复数类型： complex64、complex128
//字符串类型： string
//字符类型： byte、rune

//73 interface{}是可以指向任意对象的Any类型（）
//参考答案：对的

//74 下面关于文件操作的代码可能触发异常（）
//file, err := os.Open("test.go")
//defer file.Close()
//if err != nil {
//fmt.Println("open file failed:", err)
//return
//}
//...
//参考答案：对的 defer注册应在错误判断之后进行
// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.

//75 Golang不支持自动垃圾回收（）
//参考答案：错的，支持的
// 对于实时性要求比较高的程序，一定要关注GC问题。因为golang的GC非常稚嫩，与java相比还差的很远

//76 Golang支持反射，反射最常见的使用场景是做对象的序列化（）
//参考答案：对的

//77 Golang可以复用C/C++的模块，这个功能叫Cgo（）
//参考答案：对的

//78 下面代码中两个斜点之间的代码，比如json:"x"，作用是X字段在从结构体实例编码到JSON数据格式的时候，
//使用x作为名字，这可以看作是一种重命名的方式（）
// 字段标签属于结构体的一部分，标签不同，结构体不同
//type Position struct {
//	X int `json:"x"`
//	Y int `json:"y"`
//	Z int `json:"z"`
//}
//参考答案：对的， 输出的json{"x":1,"y":2,"z":3}

//79 通过成员变量或函数首字母的大小写来决定其作用域（）
//参考答案：对的。大写对外可见

//80 对于常量定义zero(const zero = 0.0)，zero是浮点型常量（）
//参考答案：错的， 是float64
func TestConstZero(t *testing.T) {
	const zero = 0.1
	fmt.Printf("%#v - %T", zero, zero) // 0.1 - float64
}

//
//81 对变量x的取反操作是~x（）
//参考答案：错的. go中取反操作应该是^

// 一般语言中，按位取反是：~
func TestGetReverse(t *testing.T) {
	var1 := uint(8)
	//var2 := !var1 //
	var3 := ^var1
	fmt.Printf("%08b - %T\n", var1, var1) // 8 - int
	//fmt.Printf("%#v - %T", ~var1, ~var1) // Unexpected '~'
	fmt.Printf("%08b - %T\n", var3, var3) // -9 - int
	// 00001000 - uint
	//1111111111111111111111111111111111111111111111111111111111110111 - uint
	b := true
	f := !b
	fmt.Println(b, f) // true false
}

//82 下面的程序的运行结果是xello（）
//func main() {
//	str := "hello"
//	str[0] = 'x'
//	fmt.Println(str)
//}
//参考答案：错的， str[0] = 'x' 这句话编译都不能通过，Go字符串是不可变的

//83 golang支持goto语句（）
//参考答案：对的
func TestGoto(t *testing.T) {
	x := 1
	fmt.Println("run", x) // run 1
here:
	x++
	goto here
	fmt.Println("run", x)
}

// testing 测试用例goto之后没有执行？但是并没有报错呢

//84 下面代码中的指针p为野指针，因为返回的栈内存在函数结束时会被释放（）
//type TimesMatcher struct {
//	base int
//}
//func NewTimesMatcher(base int) *TimesMatcher{
//	return &TimesMatcher{base:base}
//}
//func main() {
//	p := NewTimesMatcher(3)
//	...
//}
//参考答案：错的，NewTimesMatcher函数后不会被释放，栈逃逸技术

//85 匿名函数可以直接赋值给一个变量或者直接执行（）
//参考答案：对的
func TestAnonyFunc(t *testing.T) {
	f1 := func(a, b, c int) { fmt.Println(a, b, c) }
	f1(1, 2, 3)
}

//86 如果调用方调用了一个具有多返回值的方法，但是却不想关心其中的某个返回值，
//可以简单地用一个下划线“_”来跳过这个返回值，该下划线对应的变量叫匿名变量（）
//参考答案：对的

//87 在函数的多返回值中，如果有error或bool类型，则一般放在最后一个（）
//参考答案：对的

//88 错误是业务过程的一部分，而异常不是（）
//参考答案：对的  这是一道100%狗屎题
// 链接：https://www.nowcoder.com/questionTerminal/1ed1e4ac2dc142f980ef995b4dd973aa?orderByHotValue=1&page=1&onlyReference=false
//来源：牛客网
//
//错误指的是可能出现问题的地方出现了问题，比如打开一个文件时失败，这种情况在人们的意料之中 ；而异常指的是不应该出现问题的地方出现了问题，比如引用了空指针，这种情况在人们的意料之外。可见，错误是业务过程的一部分，而异常不是 。
//
//Golang中引入error接口类型作为错误处理的标准模式，如果函数要返回错误，则返回值类型列表中肯定包含error。error处理过程类似于C语言中的错误码，可逐层返回，直到被处理。
//
//Golang中引入两个内置函数panic和recover来触发和终止异常处理流程，同时引入关键字defer来延迟执行defer后面的函数。
//一直等到包含defer语句的函数执行完毕时，延迟函数（defer后的函数）才会被执行，而不管包含defer语句的函数是通过return的正常结束，还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。
//当程序运行时，如果遇到引用空指针、下标越界或显式调用panic函数等情况，则先触发panic函数的执行，然后调用延迟函数。调用者继续传递panic，因此该过程一直在调用栈中重复发生：函数停止执行，调用延迟执行函数等。如果一路在延迟函数中没有recover函数的调用，则会到达该携程的起点，该携程结束，然后终止其他所有携程，包括主携程（类似于C语言中的主线程，该携程ID为1）。
//
//错误和异常从Golang机制上讲，就是error和panic的区别。很多其他语言也一样，比如C++/Java，没有error但有errno，没有panic但有throw。
//
//Golang错误和异常是可以互相转换的：
//
//错误转异常，比如程序逻辑上尝试请求某个URL，最多尝试三次，尝试三次的过程中请求失败是错误，尝试完第三次还不成功的话，失败就被提升为异常了。
//异常转错误，比如panic触发的异常被recover恢复后，将返回值中error类型的变量进行赋值，以便上层函数继续走错误处理流程。

//89 函数执行时，如果由于panic导致了异常，则延迟函数不会执行（）
//参考答案：错的，仍然会执行

//90 当程序运行时，如果遇到引用空指针、下标越界或显式调用panic函数等情况，
//则先触发panic函数的执行，然后调用延迟函数。
//调用者继续传递panic，因此该过程一直在调用栈中重复发生：
//函数停止执行，调用延迟执行函数。如果一路在延迟函数中没有recover函数的调用，则会到达该携程的起点，
//该携程结束，然后终止其他所有携程，其他携程的终止过程也是重复发生：函数停止执行，调用延迟执行函数（）
//参考答案：错的，延迟函数会先压栈，然后出栈执行。

//91 同级文件的包名不允许有多个（）
//参考答案：对的    当你看到main包，你的良心不痛么。辣鸡题目
// 同级目录下，只能有一个包名，但包名可与该级目录名不一致
// 一个文件夹下只能有一个包，可以多个.go文件，但这些文件必须属于同一个包。

//92 可以给任意类型添加相应的方法（）
//参考答案：错的   内置类型就不行

//93 golang虽然没有显式的提供继承语法，但是通过匿名组合实现了继承（）
//参考答案：对的   golang面向对象不充分，组合来实现继承

//94 使用for range迭代map时每次迭代的顺序可能不一样，因为map的迭代是随机的（）
//参考答案：对的

//95 switch后面可以不跟表达式（）
//参考答案：对的
func TestSwitch(t *testing.T) {
	switch {
	}
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}
}

//96 结构体在序列化时非导出变量（以小写字母开头的变量名）不会被encode，
//因此在decode时这些非导出变量的值为其类型的零值（）
//参考答案：对的
type Person struct {
	Name string
	age  int
}

func TestSerial(t *testing.T) {
	p := Person{Name: "youyifeng", age: 28}
	m1, err := json.Marshal(p)
	if err != nil {
		fmt.Println("marshal person err", err)
		return
	}
	buff := Person{}
	err = json.Unmarshal(m1, &buff)
	if err != nil {
		fmt.Println("unmarshal person err", err)
		return
	}
	fmt.Printf("%#v - %T", buff, buff)
	// main.Person{Name:"youyifeng", age:0} - main.Person
}

//97 golang中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名（）
//参考答案：对的

//98 当函数deferDemo返回失败时，并不能destroy已create成功的资源（）
//func deferDemo() error {
//	err := createResource1()
//	if err != nil {
//		return ERR_CREATE_RESOURCE1_FAILED
//	}
//	defer func() {
//		if err != nil {
//			destroyResource1()
//		}
//	}()
//
//	err = createResource2()
//	if err != nil {
//		return ERR_CREATE_RESOURCE2_FAILED
//	}
//	defer func() {
//		if err != nil {
//			destroyResource2()
//		}
//	}()
//
//	err = createResource3()
//	if err != nil {
//		return ERR_CREATE_RESOURCE3_FAILED
//	}
//	return nil
//}
//参考答案：错的，有defer啊
//链接：https://www.nowcoder.com/questionTerminal/6501a1f6a0804ab59228cb05b09fce84?source=relative
//来源：牛客网
//
//函数返回失败有3种情况：
//一是第一次分配资源失败，直接返回，这时并没有分配成功的资源；
//二是第一次分配资源成功，第二次分配资源失败，函数返回，第二次和第三次的资源都未成功分配，此时err不为nil，第一次分配成功的资源通过defer释放；
//三是第一二次资源分配成功，第三次资源分配失败，函数返回，第一二次分配成功的资源通过defer释放；
//如果第三次资源分配也成功了，则函数不会返回失败。


//99 channel本身必然是同时支持读写的，所以不存在单向channel（）
//参考答案：错的  chan作为参数时可以限定

//100 import后面的最后一个元素是包名（）
//参考答案：错的  import后面跟的是包的路径，而不是包名。

func main() {

}
