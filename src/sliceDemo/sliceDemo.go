package sliceDemo

import (
	"fmt"
)

//切片的内存布局,简单模拟
type slice struct {
	ptr *[5]int //指向数组的指针
	len int
	cap int
}

//切片的快速实现
func make1(s slice, cap int) slice {
	s.ptr = new([5]int)
	s.cap = cap
	s.len = 0
	return s
}
func SliceDemp() {
	//	切片的初始化有两种
	//1.数组初始化
	//2.make方法快速初始化
	//	var slice1 int  = array[start:end]  从数组start开始，end结束，不包含end；start，end可以根据选择的限定省略
	//	var slice1 int  = array[0:len(array)]  等同于 array[:]
	var slice []int
	var array [5]int = [...]int{1, 2, 3, 4, 5}
	array[0] = 10

	slice = array[:]
	fmt.Println(slice)
	slice[1] = 100

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice1 := slice[0:3]
	fmt.Println(slice1)
	slice1[0] = 423
	fmt.Println(slice1)
	fmt.Println(slice)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	//从上面的例子可以看出来，slice是array的引用，所以虽然len不同，但是cap都会一样；因为就是引用的array，array的cap是多少就是多少

	var s2 = make([]int, len(slice))
	copy(slice1, slice)
	s2[0] = 200
	fmt.Println(s2)

	s2 = append(s2, 300)
	fmt.Println(slice1)
	fmt.Println(s2)

}
