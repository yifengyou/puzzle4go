package main

import "fmt"

//以下代码能编译过去吗？为什么？
type People interface {
	Speak(string) string
}
type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	var peo People = &Stduent{}

	//var peo People = Stduent{}
	// Cannot use 'Stduent{}' (type Stduent) as the type People Type does not
	// implement 'People' as the 'Speak' method has a pointer receiver

	think := "bitch"
	fmt.Println(peo.Speak(think))
}

// You are a good boy


//考点：golang的方法集
//
//解答：编译不通过！
//只是Stduent的指针类型实现了Speak这个方法， 所以，报错，编译不通过了。
