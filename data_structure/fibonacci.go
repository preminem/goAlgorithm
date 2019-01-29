package dataStruct

//闭包的调用理念
//1.会被外界调用一次来初始化
//2.返回一个函数f给调用方
//3.调用方接下去只要重复调用f就可以达到目的

//闭包的编程套路
//1.在闭包函数中且在返回函数的体外，定义每下一次计算需要记忆的变量，如left和right
//2.在返回函数的体中去实现对于记忆变量的每次更新
//3.在返回函数的体中返回调用方希望得到的结果

func fibonacci() func() int {
	x1, x2 := 0, 1
	sum := 0
	return func() int {
		sum = x1 + x2
		x1 = x2
		x2 = sum
		return sum
	}
}

func StartFibonacci(numbers int) {
	f := fibonacci()
	for i := 0; i < numbers; i++ {
		fmt.Print(f(), ", ")
	}
	fmt.Println("...")
}
