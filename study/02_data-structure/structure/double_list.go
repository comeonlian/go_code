package structure

import (
	"fmt"
	"sync"
)

// DoubleList 双向链表
type DoubleList struct {
	head, tail *ListNode
	len        int
	lock       sync.Mutex
}

func (d *DoubleList) Length() (v int) {
	v = d.len
	return
}

func (d *DoubleList) AddFromHead(p string) {
	d.lock.Lock()
	defer d.lock.Unlock()

	node := new(ListNode)
	node.value = p

	if d.head == nil || d.tail == nil{
		d.head = node
		d.tail = node
	} else {
		curNode := d.head

		curNode.prev = node
		node.next = curNode

		d.head = node
	}

	d.len += 1
}

func (d *DoubleList) AddFromTail(p string) {
	d.lock.Lock()
	defer d.lock.Unlock()

	node := new(ListNode)
	node.value = p

	if d.tail == nil || d.head == nil {
		d.tail = node
		d.head = node
	} else {
		curNode := d.tail

		curNode.next = node
		node.prev = curNode

		d.tail = node
	}

	d.len += 1
}

func (d *DoubleList) IndexFromHead(p int) (v *ListNode) {
	if d.len == 0 {
		return
	}

	curNode := d.head
	for i := 1; i < p; i++ {
		curNode = curNode.next
	}

	v = curNode
	return
}

func (d *DoubleList) RemoveFromHead(p int) (v *ListNode) {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.len == 0 || p > d.len {
		return
	}

	if d.len == 1 {
		d.head = nil
		d.tail = nil
	} else {
		cur := d.head

		for i := 1; i <= p; i++ {
			cur = cur.next
		}
		if !cur.HasPrev() {
			d.head = cur.next
		} else {
			cur.prev.next = cur.next
		}
	}

	d.len -= 1
	return
}

func (d *DoubleList) Print() {
	if d.len == 0 {
		fmt.Println("double list is empty")
	}

	curNode := d.head
	for curNode!= nil {
		fmt.Printf("%v ", curNode.value)
		curNode = curNode.next
	}
}