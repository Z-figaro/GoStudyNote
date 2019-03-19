package closureDemo

import "fmt"

func ClosureDemo() {
	//闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

//闭包
func adder() func(int) int {

	sum := 0
	innerfunc := func(x int) int {
		sum += x
		return sum
	}

	return innerfunc
}
