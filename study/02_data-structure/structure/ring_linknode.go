package structure

type RingLinkNode struct {
	Data int64
	Prev, Next *RingLinkNode
}
