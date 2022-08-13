package structure

import "sync"

type LinkQueueNode struct {
	value string
	next  *LinkQueueNode
}

type LinkQueue struct {
	root *LinkQueueNode
	size int
	lock sync.Mutex
}

func (queue *LinkQueue) Push(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	node := new(LinkQueueNode)
	node.value = v

	if queue.root == nil {
		queue.root = node
	} else {
		curNode := queue.root
		for curNode.next != nil {
			curNode = curNode.next
		}
		curNode.next = node
	}

	queue.size += 1
}

func (queue *LinkQueue) Pop() (v string) {
	if queue.size == 0 || queue.root == nil{
		return
	}

	queue.lock.Lock()
	defer queue.lock.Unlock()

	curNode := queue.root
	nextNode := curNode.next
	queue.root = nextNode

	v = curNode.value
	return
}
