package problem0027

/*
Topic:
移除元素
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

Example:
示例 1:
给定 nums = [3,2,2,3], val = 3,
函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,1,2,2,3,0,4,2], val = 2,
函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
注意这五个元素可为任意顺序。
你不需要考虑数组中超出新长度后面的元素。
 */

 /*
方法一：双指针
思路
既然问题要求我们就地删除给定值的所有元素，我们就必须用 O(1)O(1) 的额外空间来处理它。如何解决？我们可以保留两个指针 ii 和 jj，其中 ii 是慢指针，jj 是快指针。

算法
当 nums[j]与给定的值相等时，递增j以跳过该元素。只要 nums[j]！=val
我们就复制 nums[j]到nums[i] 并同时递增两个索引。重复这一过程，直到j到达数组的末尾，该数组的新长度为i。
该解法与删除排序数组中的重复项的解法十分相似。
 */
func removeElementA(nums []int, val int) int {
	i:=0
	for j:=0;j<len(nums);j++{
		if nums[j] != val{
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

/*
方法二：双指针 —— 当要删除的元素很少时
思路
现在考虑数组包含很少的要删除的元素的情况。例如，num=[1，2，3，5，4]，Val=4。之前的算法会对前四个元素做不必要的复制操作。
另一个例子是 num=[4，1，2，3，5]，Val=4。似乎没有必要将 [1，2，3，5]这几个元素左移一步，因为问题描述中提到元素的顺序可以更改。

算法
当我们遇到 nums[i] = val时，我们可以将当前元素与最后一个元素进行交换，并释放最后一个元素。这实际上使数组的大小减少了1。
请注意，被交换的最后一个元素可能是您想要移除的值。但是不要担心，在下一次迭代中，我们仍然会检查这个元素。

 */

func removeElementB(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for{
		if i>=j{
			break
		}
		if nums[i] == val {
			nums[i] = nums[j]
			j--
		}else{
			i++
		}
	}
	return j+1
}

func removeElementC(nums []int, val int) int {
	// j指向最后一个不为val的位置
	// i指向第一个值为val的位置
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i] != val {
			i++
		}

		for j >= 0 && nums[j] == val {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	return i
}