package main

//关于channel，下面语法正确的是（）
//A. var ch chan int
//B. ch := make(chan int)
//C. <- ch
//D. ch <-

//AB就是初始化一个channel
//C 从channel中取值
//D是要赋值的，值在哪呢？ 这是搞啥捏？大哥

// chan可以没有接收者，但是发送者不能无中生有
// 钱太多了可以捐，可以万能，但是，没钱，万万不能

func main() {
	ch := make(chan int)
	//ch <-  // <expression> expected, got '}'
	<-ch
}

//参考答案：ABC
