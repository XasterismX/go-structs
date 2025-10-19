package BinarySearchTree

import "errors"

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

type BinarySearchTree struct {
	root *Node
	size int
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}
func (bt *BinarySearchTree) Size() int {
	return bt.size
}
func (bt *BinarySearchTree) Insert(value int) {
	if bt.size == 0 {
		bt.root = &Node{value, nil, nil, nil}
		bt.size++
		return
	}
	current := bt.root
	for i := 0; i < bt.size; i++ {
		if current.value > value {
			if current.left == nil {
				current.left = &Node{value, nil, nil, current}
				bt.size++
				return
			} else {
				current = current.left
				continue
			}
		}
		if current.value < value {
			if current.right == nil {
				current.right = &Node{value, nil, nil, current}
				bt.size++
				return
			} else {
				current = current.right
				continue
			}
		}

	}

}
func (bt *BinarySearchTree) Find(value int) (*Node, error) {

	current := bt.root
	for i := 0; i < bt.size; i++ {
		if current.value == value {
			return current, nil
		}
		if current.value > value {
			current = current.left
			continue
		}
		if current.value < value {
			current = current.right
			continue
		}

	}
	return nil, errors.New("not found")
}
