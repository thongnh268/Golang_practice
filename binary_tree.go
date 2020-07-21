package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

//printInOrder
func (t *TreeNode) printInOrder() {
	if t == nil {
		return
	}
	t.left.printInOrder()
	fmt.Println(t.val)
	t.right.printInOrder()
}

//printPreOrder
func (t *TreeNode) printPreOrder() {
	if t == nil {
		return
	}
	fmt.Println(t.val)
	t.left.printPreOrder()
	t.right.printPreOrder()
}

//printPostOrder
func (t *TreeNode) printPostOrder() {
	if t == nil {
		return
	}
	t.left.printPostOrder()
	t.right.printPreOrder()
	fmt.Println(t.val)
}

//Insert a new node to the binary tree
func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.val == value {
		return errors.New("This node value already exists")
	}

	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{
				val: value,
			}
			return nil
		}
		return t.right.Insert(value)
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{
				val: value,
			}
			return nil
		}
		return t.left.Insert(value)
	}
	return nil
}

//Find the treenode for the given node val

func (t *TreeNode) Find(value int) (TreeNode, bool) {
	if t == nil {
		return TreeNode{}, false
	}
	switch {
	case value == t.val:
		return *t, true
	case value < t.val:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
	}
}

//Delete removes the Item with value from the tree
func (t *TreeNode) Delete(value int) {
	t.remove(value)
}

func (t *TreeNode) remove(value int) *TreeNode {
	if t == nil {
		return nil
	}

	if value < t.val {
		t.left = t.left.remove(value)
		return t
	}

	if value > t.val {
		t.right = t.right.remove(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.right == nil {
		t = t.left
		return t
	}

	if t.left == nil {
		t = t.right
		return t
	}

	smallestValOnRight := t.right
	for {
		//find smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.remove(t.val)
	return t
}

func (t *TreeNode) findMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.findMax()
}

func (t *TreeNode) findMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.findMin()
}

func (t *TreeNode) CountLeaves() int {
	if t == nil {
		return 0
	} else {
		ld := t.left.CountLeaves()
		rd := t.right.CountLeaves()
		return 1 + ld + rd
	}
}

func (t *TreeNode) Depth() int {
	if t == nil {
		return 0
	} else {
		ld := t.left.Depth()
		rd := t.right.Depth()
		switch {
		case ld > rd:
			return 1 + ld
		default:
			return 1 + rd
		}
	}
}

func main() {
	t := &TreeNode{val: 4}

	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(8)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)

	t.Find(11)

	t.Delete(5)
	t.Delete(7)
	t.printInOrder()
	fmt.Println("")
	fmt.Println(t.CountLeaves())
	fmt.Println(t.Depth())
	fmt.Println("min is %d", t.findMin())
	fmt.Println("max is %d", t.findMax())
}
