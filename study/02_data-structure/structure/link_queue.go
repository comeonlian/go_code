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

	if queue.root == nil {
		queue.root = new(LinkQueueNode)
		queue.root.value = v
	} else {
		
	}

	queue.size += 1
}

func (queue *LinkQueue) Pop() (v string) {

}
