package main

import (
	"errors"
	"fmt"
)

//从切片中删除一个元素，下面的算法实现正确的是（）
//A.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			if i== len(*s) - 1 {
//				*s = (*s)[:i]
//			}else {                  // 如果只有一个元素
//				*s = append((*s)[:i],(*s)[i + 2:]...)
//			}
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}

//B.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			*s = append((*s)[:i],(*s)[i + 1:])
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}

//C.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			delete(*s, v) // delete 用于字典
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}

//D.
//func (s *Slice)Remove(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			*s = append((*s)[:i],(*s)[i + 1:]...)
//			return nil
//		}
//	}
//	return ERR_ELEM_NT_EXIST
//}
//参考答案：D

// 整个逻辑就是，找到你要删除的元素，然后把它前面的和后面的组合起来
//  delete函数**只能**用于删除map中对应key的键值对，如果map中不存在该key，则什么也不做
// 此处append有一个坑要注意 *s = append((*s)[:i],(*s)[i + 1:]...) 理论上越界，但实际并没有

type Slice []byte

func isEqual(value interface{}, v interface{}) bool {
	fmt.Println(value, v, value == v)
	fmt.Printf("%T - %T\n", value, v)
	if value == v {
		fmt.Println("FIND EQUAL")
		return true
	}
	return false
}
func (s *Slice) Remove(value interface{}) error {
	for i, v := range *s {
		fmt.Printf("index:%d value:%d\n", i, v)
		if isEqual(value, int(v)) {
			fmt.Println("i+1:", i+1)
			fmt.Println("s:", (*s)[i+1:])
			//fmt.Println("s:", (*s)[i+2:]) // runtime error: slice bounds out of range [4:3]
			*s = append((*s)[:i], (*s)[i+1:]...)
			fmt.Println(*s)
			return nil
		}
	}
	return errors.New("Not found!")
}

func (s *Slice)Remove2(value interface{}) error {
	for i, v := range *s {
		if isEqual(value, v) {
			if i== len(*s) - 1 {
				*s = (*s)[:i]
			}else {
				*s = append((*s)[:i],(*s)[i + 2:]...)
			}
			return nil
		}
	}
	return errors.New("Not found!")
}


//func (s *Slice)Remove3(value interface{}) error {
//	for i, v := range *s {
//		if isEqual(value, v) {
//			*s = append((*s)[:i],(*s)[i + 1:])
//			// cannot use (*s)[i + 1:] (type Slice) as type byte in append
//			return nil
//		}
//	}
//	return errors.New("Not found!")
//}


func main() {
	ss := []byte{0, 1, 2}
	//fmt.Println(ss[len(ss)]) // index out of range [3] with length 3
	sp := &ss
	fmt.Println((*sp)[0])
	fmt.Println((*sp)[1])
	fmt.Println((*sp)[2])
	//fmt.Println((*sp)[3]) // index out of range [3] with length 3
	s2 := append(*sp, (*sp)[3:]...)
	fmt.Println(s2)
	fmt.Println("=========================")

	s := Slice{3}
	fmt.Println(s)
	s.Remove2(3)
	fmt.Println(s)
}
