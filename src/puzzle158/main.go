package main

import "fmt"

type People interface {
}
type Student struct{}

type ShowPerson interface {
	Show()
}

func main() {
	var p People
	fmt.Printf("type:%T value:%v p==nil:%v\n", p, p, p == nil)
	fmt.Printf("%#v - %T\n", p, p)
	var i interface{}
	fmt.Println(p == i)
	var sp ShowPerson
	fmt.Println(sp == i)

	var s *Student
	p = s
	fmt.Printf("type:%T value:%v p==nil:%v\n", p, p, p == nil)
	fmt.Printf("%#v - %T\n", p, p)
	//type:<nil> value:<nil> p==nil:true
	//<nil> - <nil>
	//true
	//true
	//type:*main.Student value:<nil> p==nil:false
	//(*main.Student)(nil) - *main.Student
}
