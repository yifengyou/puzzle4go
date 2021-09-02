package main

import (
	"fmt"
	"testing"
	"time"
)

//51 声明一个整型变量i__________
//参考答案：var i int

//52 声明一个含有10个元素的整型数组a__________
//参考答案：var a [10]int

//53 声明一个整型数组切片s__________
//参考答案：var s []int

//54 声明一个整型指针变量p__________
//参考答案：var p *int

//55 声明一个key为字符串型value为整型的map变量m__________
//参考答案：var m map[string]int

//56 声明一个入参和返回值均为整型的函数变量f__________
//参考答案：var f func(a int) int

//57 声明一个只用于读取int数据的单向channel变量ch__________
//参考答案：var ch <-chan int

//58 假设源文件的命名为slice.go，则测试文件的命名为__________
//参考答案：slice_test.go

//59 go test要求测试函数的前缀必须命名为__________
//参考答案：Test    例如：TestHelloWorld

//60 下面的程序的运行结果是__________
//for i := 0; i < 5; i++ {
//defer fmt.Printf("%d ", i)
//}
//参考答案：4 3 2 1 0
func TestFor(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) // 4 3 2 1 0
	}
	// defer先注册后执行，栈
	// defer会拷贝参数，否则都是4
}

//61 下面的程序的运行结果是__________
//func main() {
//	x := 1
//	{
//		x := 2
//		fmt.Print(x)
//	}
//	fmt.Println(x)
//}
//参考答案：21   就近原则，谁离的近，就找谁，和找对象一样，你工作和哪个小姐姐，小哥哥经常接触，就最有机会

//62 下面的程序的运行结果是__________
//func main() {
//	strs := []string{"one", "two", "three"}
//
//	for _, s := range strs {
//		go func() {
//			time.Sleep(1 * time.Second)
//			fmt.Printf("%s ", s)
//		}()
//	}
//	time.Sleep(3 * time.Second)
//}
//参考答案：three three three   闭包问题，当函数执行的时候，s的值都是three

func TestRange(t *testing.T) {
	strs := []string{"one", "two", "three"}

	for _, s := range strs {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s ", s)
		}()
	}
	time.Sleep(3 * time.Second)
}

// three three three
// 此处协程中引用的变量s，不是值拷贝，而是地址闭包了。
// 而go语句与defer有共性，其实都是先注册后执行。如果协程中没有sleep，结果可能五花八门

//63 下面的程序的运行结果是__________
//func main() {
//	x := []string{"a", "b", "c"}
//	for v := range x {
//		fmt.Print(v)
//	}
//}+
//参考答案：012
func TestStr(t *testing.T) {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Printf("%#v,%T\n", v, v)
	}
	for v := range x {
		fmt.Printf("%#v,%T\n", x[v], x[v])
	}
	for key, value := range x {
		fmt.Printf("%#v,%#v\n", key, value)
	}
}

// range 一般这么写 key,value := range X
// 但如果只有一个则为从0开始的序列   num := range X

//64 下面的程序的运行结果是__________
//func main() {
//	x := []string{"a", "b", "c"}
//	for _, v := range x {
//		fmt.Print(v)
//	}
//}
//参考答案：abc

//65 下面的程序的运行结果是__________
//func main() {
//	i := 1
//	j := 2
//	i, j = j, i
//	fmt.Printf("%d%d\n", i, j)
//}
//参考答案：21

//66 下面的程序的运行结果是__________
//func incr(p *int) int {
//	*p++
//	return *p
//}
//func main() {
//	v := 1
//	incr(&v)
//	fmt.Println(v)
//}
//参考答案：2

//67 启动一个goroutine的关键字是__________
//参考答案：go

//68 下面的程序的运行结果是__________
//type Slice []int
//func NewSlice() Slice {
//	return make(Slice, 0)
//}
//func (s* Slice) Add(elem int) *Slice {
//	*s = append(*s, elem)
//	fmt.Print(elem)
//	return s
//}
//func main() {
//	s := NewSlice()
//	defer s.Add(1).Add(2)
//	s.Add(3)
//}
//参考答案：132

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func TestSlice(t *testing.T) {
	s := NewSlice()
	defer s.Add(1).Add(2)
	fmt.Printf("======")
	s.Add(3)
}
// 1======32
// defer注册函数，最终只会有一个函数待执行，前序都在注册前执行完成，这是个不大不小的坑