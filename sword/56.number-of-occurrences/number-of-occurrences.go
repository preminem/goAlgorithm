package problem56

import (
	"unsafe"
	"errors"
)

/*
题目一：数组中只出现一次的两个数字
一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度O(n)，空间复杂度O(1)。

分析：
使用异或运算的一个性质：任何一个数字异或它自己都等于0；那么如果从头到尾依次异或数组中的每个数，最终即是只出现依次的数字；
题中需要寻找两个只出现一次的数字，那么我们如果将原数组分成两个子数组，每个子数组中只包含一个数字，其他数字都是成对出现。
1.首先对所有元素进行异或，得到的结果是两个只出现一次的数字异或的结果，因为两个数字不一样，所以异或的结果中肯定至少有一位1；
2.我们以第一个1的位置是否出现1为划分，将数组分为两个子数组；
3.分别对每个子数组进行元素异或，则可以得到两个只出现一次的数字。
注意：不必进行子数组的分拆和组合，只要在第二次遍历整个数组的过程中，对每个元素进行一次属于哪个子数组的判断操作即可。


 */
func findNumbersAppearOnce(data []int) (int, int) {
	if nil == data || len(data) < 2 {
		return 0, 0
	}

	var result int
	for i := 0; i < len(data); i++ {
		result ^= data[i]
	}
	indexOf1 := findFirstBitIs1(result)
	var num1, num2 int
	for j := 0; j < len(data); j++ {
		if isBit1(data[j], indexOf1) {
			num1 ^= data[j]
		} else {
			num2 ^= data[j]
		}
	}
	return num1, num2
}

func findFirstBitIs1(num int) uint {
	var indexBit uint
	//和1做与运算可以判断最后一位是不是1
	for (0 == num&1) && (indexBit < uint(unsafe.Sizeof(int(0))*8)) {
		num = num >> 1
		indexBit++
	}
	return indexBit
}

func isBit1(num int, indexBit uint) bool {
	num = num >> indexBit
	if 1 == num&1 {
		return true
	}
	return false
}

/*
题目二：数组中唯一只出现一次的数字
在一个数组中除一个数字只出现一次之外，其他数字都出现了三次。请找出哪个只出现一次的数字。

分析：
除目标数字外其他数字都出现三次，那么其二进制表示的每一位（0或者1）也出现三次；
所有数字每一位均相加，每一位的和均能被3整除，不能被三整除的那一位即为目标数该位为1的情况
 */
func findNumberAppearOnce2(numbers []int) int {
	if nil == numbers || 0 >= len(numbers) {
		panic(errors.New("invalid input"))
	}

	bitSum := make([]int, 32)
	for i := 0; i < len(numbers); i++ {
		bitMask := 1
		for j := 31; j >= 0; j-- {
			bit := numbers[i] & bitMask
			if 0 != bit {
				bitSum[j] += 1
			}
			bitMask = bitMask << 1
		}
	}

	var result int
	for i := 0; i < 32; i++ {
		result = result << 1
		result += bitSum[i] % 3
	}
	return result
}