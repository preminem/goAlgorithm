package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Deque(t *testing.T) {
	ast := assert.New(t)

	d := NewDeque()
	ast.True(d.IsEmpty(), "检查新建的 q 是否为空")

	start, end := 0, 100

	for i := start; i < end; i++ {
		d.Push(i)
		ast.Equal(i-start+1, d.Len(), "Push 后检查 d 的长度。")
	}

	for i := end - 1; i >= start; i-- {
		ast.Equal(i, d.Pop(), "从 d 中 pop 出数来。")
	}

	ast.True(d.IsEmpty(), "检查 Pop 完毕后的 s 是否为空")

	for i := start; i < end; i++ {
		d.UnShift(i)
		ast.Equal(i-start+1, d.Len(), "UnShift 后检查 d 的长度。")
	}

	for i := end - 1; i >= start; i-- {
		ast.Equal(i, d.Shift(), "从 d 中 shift 出数来。")
	}

	ast.True(d.IsEmpty(), "检查 Shift 完毕后的 d 是否为空")
}
