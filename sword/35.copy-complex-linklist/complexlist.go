package problem35

import "fmt"

type ComplexListNode struct {
	value   int
	next    *ComplexListNode
	sibling *ComplexListNode
}

func CreateNode(value int) *ComplexListNode {
	node := ComplexListNode{}

	node.value = value
	node.next = nil
	node.sibling = nil
	return &node
}

func BuildNodes(node, next, sibling *ComplexListNode) {
	if node != nil {
		node.next = next
		node.sibling = sibling
	}
}

func PrintList(head *ComplexListNode) {
	node := head

	for node != nil {
		fmt.Printf("The value of this node is: %d.\n", node.value)
		if node.sibling != nil {
			fmt.Printf("The value of its sibling is: %d.\n", node.sibling.value)
		}
		node = node.next
	}
	fmt.Print("\n")
}
