package repeat_number

/*
两个从小到大排好序的int数组(没有重复元素)。找出所有在这两个数组中都出现过的数。

*/
func RepeatNumber(a, b []int) []int {
	res := make([]int, 0)
	i, j := 0, 0
	for {
		if i >= len(a) || j >= len(b) {
			break
		}
		if a[i] == b[j] {
			res = append(res, a[i])
			i++
			j++
		} else if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}
	return res
}
