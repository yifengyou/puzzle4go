package main

//关于switch语句，下面说法正确的有（）
//A. 条件表达式必须为常量或者整数
//B. 单个case中，可以出现多个结果选项。但是不能重复
//C. 需要用break来明确退出一个case
//D. 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case

//switch var1 {
//case val1:
//...
//case val2:
//...
//default:
//...
//}

//A 必须滴？ 是不是这种答案直接否定了，bool
//C break?  什么鬼，是不是记乱了？C中是要break

func main() {
	var r int
	switch r {
	//case 1,2,2: //Duplicate case 2
	//	break
	case 3, 4:
		break
	}
	b := true
	switch b {
	case true:
		break
	case false:
		break
	}
}

//参考答案：BD
