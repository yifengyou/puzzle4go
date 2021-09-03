package main

import (
	"fmt"
	"time"
)

//golang 两个协成 交替 打印1-100的奇数偶数


var POOL = 100


func groutine1(p chan int) {
	for i := 1; i <= POOL; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println("groutine-1:", i)
		}
	}
}
func groutine2(p chan int) {
	for i := 1; i <= POOL; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println("groutine-2:", i)
		}
	}
}
func main() {
	msg := make(chan int)
	go groutine1(msg)
	go groutine2(msg)
	time.Sleep(time.Second * 3)
}

// 考点

// 无缓冲chan的基本使用，读端阻塞除非可读，写端阻塞除非可写