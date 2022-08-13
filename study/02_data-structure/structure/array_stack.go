package structure

import "sync"

type ArrayStack struct {
	array []string   // 底层切片元素
	size  int        // 栈的元素数量
	lock  sync.Mutex // 为了并发安全使用锁
}

func (stack *ArrayStack) Push(val string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.array = append(stack.array, val)
	stack.size += 1
}

func (stack *ArrayStack) Pop() (val string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		return
	}

	val = stack.array[stack.size-1]

	newArr := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArr[i] = stack.array[i]
	}
	stack.array = newArr
	stack.size = stack.size - 1
	return
}

func (stack *ArrayStack) Peek() (val string) {
	if stack.size == 0 {
		return
	}
	val = stack.array[stack.size-1]
	return
}

func (stack *ArrayStack) Size() (size int) {
	return stack.size
}

func (stack *ArrayStack) IsEmpty() (isEmpty bool) {
	return stack.size == 0
}
