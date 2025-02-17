package puzzle101

//切片与数组对比
//
//(1). 数组
// 1.数组是具有固定长度且拥有零个或者多个相同数据类型元素的序列。
// 2.数组的长度是数组类型的一部分，所以[3]int 和 [4]int 是两种不同的数组类型。
// 3.数组需要指定大小，不指定也会根据初始化的自动推算出大小，不可改变 ;
// 4.数组是值传递，不管数组大小如何，都是拷贝
// 5.数组是内置(build-in)类型,是一组同类型数据的集合，它是值类型，通过从0开始的下标索引访问元素值。
//   在初始化后长度是固定的，无法修改其长度。当作为方法的参数传入时将复制一份数组而不是引用同一指针。
//   数组的长度也是其类型的一部分，通过内置函数len(array)获取其长度。

//(2). 切片
// 1.切片表示一个拥有相同类型元素的可变长度的序列。
// 2.切片是一种轻量级的数据结构，它有三个属性：指针、长度和容量。
// 3.切块结构如下：
	//type Slice struct {
	//	ptr unsafe.Pointer
	//	len int
	//	cap int
	//}

// 4.切片不需要指定大小;
// 5.切片是地址传递;
// 6.切片可以通过数组来初始化，也可以通过内置函数make()初始化.
//   初始化时len=cap,在追加元素时如果容量cap不足时将按len的2倍扩容；

//(3). 关系
//一个底层数组可以对应多个slice，这些slice可以引用数组的任何位置，彼此之间的元素还可以重叠。

func main(){

}