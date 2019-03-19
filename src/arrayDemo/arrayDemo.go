package arrayDemo

import "fmt"

//1.数组必须是定长的
func ArrayDemo() {

	var a [10]int

	a[0] = 100

	//两种遍历方式
	for i := 0; i < len(a); i++ {

		fmt.Println(a[i])
	}

	for index, val := range a {

		fmt.Printf("a[%d] = %d\n", index, val)
	}

}

// 数组是值类型的证明
func ArrayTypeCheck() {

	var a [5]int
	a[0] = 10
	b := a
	b[1] = 100
	//值类型的特点是改变副本的值，不会影响a本身的值
	fmt.Printf("a = %d b = %d", a, b)
}

//不使用递归实现斐波拉切数列，显示前100个

func Fab(n int) {

	var a []int
	a = make([]int, n)

	a[0] = 1
	a[1] = 1

	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]

	}
	for _, value := range a {

		fmt.Println(value)

	}
}

func ArrayInit() {
	//数组初始化的几种方式
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{1, 2, 3, 4, 5}
	var a3 = [...]int{1: 100, 3: 200}
	var a4 = [...]string{1: "hello", 4: "world"}

	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}

//TwoDimension
func TwoDimension() {

	var a [5][3]int = [...][3]int{{1, 2, 3}, {2, 3, 4}, {3, 234, 34}, {4, 324, 343}, {5, 23, 5432}}

	for index, value := range a {

		fmt.Printf("一维是 [%d - %d] \n", index, value)

		for index2, value2 := range value {

			fmt.Printf(" 二维是 [%d - %d] \n", index2, value2)

		}

	}
}
