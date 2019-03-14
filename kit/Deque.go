package kit

// Deque 是用于存放 int 的队列
type Deque struct {
	nums []int
}

// NewDeque 返回 *kit.Deque
func NewDeque() *Deque {
	return &Deque{nums: []int{}}
}

// Push 用来在数组末端添加项
func (d *Deque) Push(n int) {
	d.nums = append(d.nums, n)
}

// Shift 用来移除数组的第一个项
func (d *Deque) Shift() int {
	res := d.nums[0]
	d.nums = d.nums[1:]
	return res
}

// Pop 数组末端移除项
func (d *Deque) Pop() int {
	res := d.nums[len(d.nums)-1]
	d.nums = d.nums[:len(d.nums)-1]
	return res
}

//UnShift 在数组前端添加项
func (d *Deque) UnShift(n int) {
	res := []int{n}
	d.nums = append(res, d.nums...)
}

// Len 返回 d 的长度
func (d *Deque) Len() int {
	return len(d.nums)
}

// IsEmpty 反馈 d 是否为空
func (d *Deque) IsEmpty() bool {
	return d.Len() == 0
}

//队首元素
func (d *Deque) First() int {
	return d.nums[0]
}

//队尾元素
func (d *Deque) Last() int {
	return d.nums[len(d.nums)-1]
}
