package goAlgorithm

import (
	"testing"
)

func newValues() ([]int, int) {
	return []int{4, 12, 27, 37, 80, 81, 84, 85, 93, 93}, 80
}

func TestSequenceSearch(t *testing.T) {
	values, value := newValues()
	if SequenceSearch(values, value) != 4 {
		t.Errorf(`SequenceSearch(%v,%v) != 4`, values, value)
	}
}

func TestBinarySearch(t *testing.T) {
	values, value := newValues()
	if BinarySearch(values, value) != 4 {
		t.Errorf(`BinarySearch(%v,%v) != 4`, values, value)
	}
}

func TestInsertSearch(t *testing.T) {
	values, value := newValues()
	if InsertSearch(values, value) != 4 {
		t.Errorf(`InsertSearch(%v,%v) != 4`, values, value)
	}
}
