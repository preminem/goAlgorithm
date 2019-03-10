package problem50

/*
题：
第一个只出现一次的字符
在字符串中找出第一个只出现一次的字符。如输入"abaccdeff"，则输出'b'。

分析：
了解决这个问题，我们可以定义一个哈希表（外部空间），其键值（Key）是字符，而值（Value）是该字符出现的次数。
同时我们还需要从头开始扫描字符串两次：
（1）第一次扫描字符串时，每扫描到一个字符就在哈希表的对应项中把次数加1。（时间效率O(n)）
（2）第二次扫描时，每扫描到一个字符就能从哈希表中得到该字符出现的次数。这样第一个只出现一次的字符就是符合要求的输出。（时间效率O(n)）
*/

func firstNoRepeatingChar(str string) byte {
	if 0 >= len(str) {
		return 0
	}
	//建一个长度为256的数组来模拟哈希表，每个字母根据其ASCII码值作为数组的下标对应数组的一个数字，而数组中存储的是每个字符出现的次数。
	table := make([]int, 256)
	for i := 0; i < len(str); i++ {
		table[str[i]]++
	}
	for i := 0; i < len(str); i++ {
		if 1 == table[str[i]] {
			return byte(str[i])
		}
	}
	return 0
}
