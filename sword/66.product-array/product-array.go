package problem66

/*
题目：
给定一个数组A[0,1,…,n-1],请构建一个数组B[0,1,…,n-1],其中B中的元素B[i]=A[0]A[1]…A[i-1]*A[i+1]…*A[n-1]。不能使用除法。

思路标签：
分段进行

解答：
我们可以定义素C[i]=A[0]xA[1]x…xA[i-1]；D[i] = A[i+1]x…xA[n-1]；
C[i]可以自上到下计算：C[i] = C[i-1]xA[i-1];
D[i]可以自下到上计算：D[i] = D[i+1]xA[i+1]；
同时注意不满足条件时需要返回一个空数组。
*/

func Multiply(arrayA []int) []int {
	if len(arrayA) > 1 {
		arrayB := make([]int, len(arrayA))
		arrayB[0] = 1
		//arrayB先用来记录C[i]
		for i := 1; i < len(arrayA); i++ {
			arrayB[i] = arrayB[i-1] * arrayA[i-1]
		}
		//temp用来表示D[i]
		temp := 1
		for i := len(arrayA) - 2; i >= 0; i-- {
			temp = temp * arrayA[i+1]
			arrayB[i] = arrayB[i] * temp
		}
		return arrayB
	}
	return nil
}
