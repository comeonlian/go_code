package structure

// ListNode 链表节点
type ListNode struct {
	prev, next *ListNode
	value      string
}

func (l *ListNode) GetValue() (v string) {
	v = l.value
	return
}

func (l *ListNode) GetPrev() (v *ListNode) {
	v = l.prev
	return
}

func (l *ListNode) GetNext() (v *ListNode) {
	v = l.next
	return
}

func (l *ListNode) HasPrev() (v bool) {
	v = l.prev != nil
	return
}

func (l *ListNode) HasNext() (v bool) {
	v = l.next != nil
	return
}

func (l *ListNode) IsNil() (v bool) {
	v = l == nil
	return
}

