package LinkedList

import (
	"errors"
)

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
	tail *node
	size int
}

func (list *LinkedList) Init() *LinkedList {
	list.head = &node{
		val:  0,
		next: nil,
	}
	list.size = 0
	return list
}
func (list *LinkedList) Size() int {
	return list.size
}
func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}
func (list *LinkedList) InsertAfterValue(aV, v int) error {
	newNode := &node{val: v, next: nil}
	current := list.head
	for current.next != nil {
		if current.val == aV {
			current.next = newNode
			return nil
		}
		current = current.next

	}
	return errors.New("cannot insert node, after value doesn't exist")
}
func (list *LinkedList) InsertBeforeValue(aV, v int) error {
	current := list.head
	for current.next != nil {
		if current.next.val == aV {
			newNode := &node{val: v, next: current.next}
			current.next = newNode
			return nil
		}
		current = current.next

	}
	return errors.New("cannot insert node, after value doesn't exist")
}
func (list *LinkedList) DeleteFront() bool {
	if list.head != nil {
		list.head = list.head.next
		return true
	}
	return false
}
func (list *LinkedList) DeleteLast() (bool, error) {
	if list.head == nil {
		return false, errors.New("list is empty")
	}

	if list.head.next == nil {
		list.head = nil
		return true, nil
	}

	var current *node = list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil

	return true, nil
}

func (list *LinkedList) FindByIndex(index int) (*node, error) {
	current := list.head
	if index < 0 || index >= list.size {
		return nil, errors.New("index out of range")
	}

	for c := 0; c < index; c++ {
		current = current.next
	}
	return current, nil
}
