package goAlgorithm

//HashMap木桶(数组)的个数
const BucketCount = 16

//链表结构里数据的数据类型 键值对
type KV struct {
	Key   string
	Value string
}

//链表结构
type LinkNode struct {
	//节点数据
	Data KV
	//下一个节点
	NextNode *LinkNode
}

type HashMap struct {
	//HashMap木桶
	Buckets [BucketCount]*LinkNode
}

//创建只有头结点的链表
func CreateLink() *LinkNode {
	//头结点数据为空 是为了标识这个链表还没有存储键值对
	var linkNode = &LinkNode{KV{"", ""}, nil}
	return linkNode
}

//尾插法添加节点,返回链表总长度
func (link *LinkNode) AddNode(data KV) int {
	var count = 0
	//找到当前链表尾节点
	tail := link
	for {

		count += 1
		if tail.NextNode == nil {

			break
		} else {

			tail = tail.NextNode
		}
	}
	var newNode = &LinkNode{data, nil}
	tail.NextNode = newNode
	return count + 1
}

//创建HashMap
func CreateHashMap() *HashMap {
	myMap := &HashMap{}
	//为每个元素添加一个链表对象
	for i := 0; i < BucketCount; i++ {
		myMap.Buckets[i] = CreateLink()
	}
	return myMap
}

//自定义一个简单的散列算法，它可以将不同长度的key散列成0-BucketCount的整数
func HashCode(key string) int {
	var sum = 0
	for i := 0; i < len(key); i++ {
		//字符串中每个字符都会对应一个整数
		sum += int(key[i])
	}
	return (sum % BucketCount)
}

//添加键值对
func (myMap *HashMap) AddKeyValue(key string, value string) {
	//1.将key散列成0-BucketCount的整数作为Map的数组下标
	var mapIndex = HashCode(key)
	//2.获取对应数组头结点
	var link = myMap.Buckets[mapIndex]
	//3.在此链表添加结点
	if link.Data.Key == "" && link.NextNode == nil {
		//如果当前链表只有一个节点，说明之前未有值插入  修改第一个节点的值 即未发生哈希碰撞
		link.Data.Key = key
		link.Data.Value = value
	} else {
		//发生哈希碰撞
		index := link.AddNode(KV{key, value})
	}
}

//按键取值
func (myMap *HashMap) GetValueForKey(key string) string,bool {
	//1.将key散列成0-BucketCount的整数作为Map的数组下标
	var mapIndex = HashCode(key)
	//2.获取对应数组头结点
	var link = myMap.Buckets[mapIndex]
	var value string
	//遍历找到key对应的节点
	head := link
	for {
		if head == nil{
			return nil,false
		}
		if head.Data.Key == key {
			value = head.Data.Value
			return value,true
		} else {

			head = head.NextNode
		}
	}
}

func (myMap *HashMap) DeleteValueForKey(key string) {
  var mapIndex = HashCode(key)
  //头节点指向下个节点的指针置空
  myMap.Buckets[mapIndex].NextNode = nil
}
