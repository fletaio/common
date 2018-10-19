package queue

import (
	"git.fleta.io/fleta/common/hash"
)

// LinkedQueue TODO
type LinkedQueue struct {
	Head    *LinkedItem
	Tail    *LinkedItem
	keyHash map[hash.Hash256]*LinkedItem
}

// NewLinkedQueue TODO
func NewLinkedQueue() *LinkedQueue {
	q := &LinkedQueue{
		keyHash: map[hash.Hash256]*LinkedItem{},
	}
	return q
}

// Push TODO
func (q *LinkedQueue) Push(Key hash.Hash256, item interface{}) {
	nd := &LinkedItem{
		Key:  Key,
		Item: item,
	}
	if q.Head == nil {
		q.Head = nd
		q.Tail = nd
	} else {
		nd.Prev = q.Tail
		q.Tail.Next = nd
		q.Tail = nd
	}
	q.keyHash[Key] = nd
}

// Pop TODO
func (q *LinkedQueue) Pop() interface{} {
	if q.Head == nil {
		return nil
	}
	nd := q.Head
	if nd == q.Tail {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = nd.Next
		nd.Next.Prev = nil
	}
	nd.Prev = nil
	nd.Next = nil
	delete(q.keyHash, nd.Key)
	return nd.Item
}

// Remove TODO
func (q *LinkedQueue) Remove(Key hash.Hash256) interface{} {
	if q.Head == nil {
		return nil
	}
	nd, has := q.keyHash[Key]
	if !has {
		return nil
	}
	if nd.Next != nil {
		nd.Next.Prev = nd.Prev
	}
	if nd.Prev != nil {
		nd.Prev.Next = nd.Next
	}
	if nd == q.Head {
		q.Head = nd.Next
	}
	if nd == q.Tail {
		q.Tail = nd.Prev
	}
	nd.Prev = nil
	nd.Next = nil
	return nd.Item
}

// LinkedItem TODO
type LinkedItem struct {
	Prev *LinkedItem
	Key  hash.Hash256
	Item interface{}
	Next *LinkedItem
}
