package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	val      int
	leftPtr  *BinaryNode
	rightPtr *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(val int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{
			val:      val,
			leftPtr:  nil,
			rightPtr: nil,
		}
	} else {
		t.root.insert(val)
	}
	return t
}

func (n *BinaryNode) insert(val int) {
	if n == nil {
		return
	} else if val < n.val {
		if n.leftPtr == nil {
			n.leftPtr = &BinaryNode{
				val:      val,
				leftPtr:  nil,
				rightPtr: nil,
			}
		} else {
			n.leftPtr.insert(val)
		}
	} else if val > n.val {
		if n.rightPtr == nil {
			n.rightPtr = &BinaryNode{
				val:      val,
				leftPtr:  nil,
				rightPtr: nil,
			}
		} else {
			n.rightPtr.insert(val)
		}
	}
}

// func countNodes (n *BinaryNode){
// 	if n == nil {
// 		return 0
// 	}else {
// 		int ld = countNodes (n.leftPtr);

// 	}
// }

// func (t *BinaryTree) FindMin() int {
// 	if t.root.leftPtr == nil {
// 		return t.val
// 	}
// 	return t.root.leftPtr.FindMin()
// }

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.val)
	print(w, node.leftPtr, ns+2, 'L')
	print(w, node.rightPtr, ns+2, 'R')
}
func main() {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(os.Stdout, tree.root, 0, 'M')

}
