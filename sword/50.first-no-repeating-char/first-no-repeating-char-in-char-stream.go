package problem50

/*
题：
字符流中第一个只出现一次的字符

分析：
1.因为是字符流，所以字符只能一个接着一个从字符流中读出来。
2.仍然使用哈希表来解决，这次哈希表中不记录字符出现的次数，而是记录只出现一次的字符首次出现的位置，哈希表中的元素全部初始化为-1，
如果出现一次以上则置为-2。因此，每次遍历哈希表就能找到第一个只出现一次的字符。

*/
type Container struct {
	//哈希表
	HashTable []int
	//当前从数据流中提取的字符的个数
	Index int
}

//初始化
func NewHashTable() Container {
	table := make([]int, 256)
	for i, _ := range table {
		table[i] = -1
	}
	container := Container{HashTable: table, Index: 0}
	return container
}

//从字符流中提取一个字符，并更新哈希表
func (c *Container) Insert(ch byte) {
	if c.HashTable[ch] == -1 {
		c.HashTable[ch] = c.Index
	} else if c.HashTable[ch] >= 0 {
		c.HashTable[ch] = -2
	}

	c.Index++
}

//遍历哈希表，返回第一个只出现一次的字符
func (c *Container) FirstAppearingOnce() byte {
	fstPosi := int(^uint(0) >> 1)
	var resu byte = 'a'
	for i := 0; i != 256; i++ {
		if c.HashTable[i] >= 0 {
			if c.HashTable[i] < fstPosi {
				fstPosi = c.HashTable[i]
				resu = byte(i)
			}
		}
	}
	return resu
}
