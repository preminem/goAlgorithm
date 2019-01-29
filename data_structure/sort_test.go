package dataStruct

import (
	"testing"
)

func newValues() []int {
	return []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
}

func BenchmarkChangeSort(b *testing.B) {
	values := newValues()
	for i := 0; i < b.N; i++ {
		ChangeSort(values)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	values := newValues()
	for i := 0; i < b.N; i++ {
		BubbleSort(values)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	values := newValues()
	for i := 0; i < b.N; i++ {
		InsertSort(values)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	values := newValues()
	for i := 0; i < b.N; i++ {
		QuickSort(values)
	}
}

func BenchmarkMergeSortEntrance(b *testing.B) {
	values := newValues()
	for i := 0; i < b.N; i++ {
		MergeSortEntrance(values)
	}
}

func BenchmarkShellSort(b *testing.B) {
	values := newValues()
	for i := 0; i < b.N; i++ {
		ShellSort(values)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	values := newValues()
	for i := 0; i < b.N; i++ {
		HeapSort(values)
	}
}
