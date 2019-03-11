package problem51

/*
题：
数组中的逆序对
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数。

分析：
类似于归并排序。先把数组等分成子数组，统计出子数组的逆序对数目，在统计子数组的过程中，同时对子数组进行从小到大排序。然后，递归统计前后相邻子数组之间的逆序对数目。
统计前后相邻子数组逆序对的过程如下：
使用两个指针分别指向前后子数组的末尾，然后比较指针指向的数字大小。
若前面的数字大于后面的数字，说明有逆序对，且逆序对的数目等于后面数组剩余数字的数目（因为，后半段剩余的数字都小于或等于后半段指针指向的数字）。
若前面的数字小于或等于后面的数字，说明无逆序对。每次比较完，都把较大的数字从后向前复制到辅助数组，保证辅助数组按从小到大排序，有序的辅助数组将用于求下一次更大的相邻子数组逆序对数目。
这个过程就相当于在归并排序的过程中，统计逆序对的数目。

 */

func inversePairs(data []int) int {
	if nil == data || 0 > len(data) {
		return 0
	}

	copy := make([]int, len(data))
	for id, value := range data {
		copy[id] = value
	}

	count := inversePairsCore(data, copy, 0, len(data)-1)
	return count
}

func inversePairsCore(data, copy []int, start, end int) int {
	//子数组只有一个数
	if start == end {
		copy[start] = data[start]
		return 0
	}
	length := (end - start) / 2
	//copy和data数组在这里交换位置
	//相当于每次递归后data的前后半段都是有序的
	left := inversePairsCore(copy, data, start, start+length)
	right := inversePairsCore(copy, data, start+length+1, end)
	//前半段最后一个数字的下标
	i := start + length
	//后半段最后一个数字的下标
	j := end
	//从后向前复制数字到辅助数组，保证从小到大排序
	indexCopy := end
	count := 0
	//从后向前遍历
	for i >= start && j >= start+length+1 {
		//前半段指针指向的数字大于后半段指针指向的数字
		if data[i] > data[j] {
			copy[indexCopy] = data[i]
			indexCopy--
			i--
			//逆序对的数目等于后半段剩余数字的个数
			//因为，后半段剩余的数字都小于或等于后半段指针指向的数字
			count += j - start - length
		} else {
			copy[indexCopy] = data[j]
			indexCopy--
			j--
		}
	}
	for ; i >= start; i-- {
		copy[indexCopy] = data[i]
		indexCopy--
	}
	for ; j >= start+length+1; j-- {
		copy[indexCopy] = data[j]
		indexCopy--
	}
	return left + right + count
}