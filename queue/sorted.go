package queue

import (
	"sort"
	"sync"
)

// SortedQueue TODO
type SortedQueue struct {
	sync.Mutex
	items []*sortedQueueItem
	head  int
	size  int
}

// NewSortedQueue TODO
func NewSortedQueue() *SortedQueue {
	q := &SortedQueue{
		items: make([]*sortedQueueItem, 0, 256),
	}
	return q
}

// Insert TODO
func (q *SortedQueue) Insert(value interface{}, Priority uint64) {
	q.Lock()
	defer q.Unlock()

	item := sortedQueueItemPool.Get().(*sortedQueueItem)
	item.value = value
	item.priority = Priority
	idx := sort.Search(q.size, func(i int) bool {
		return Priority < q.items[q.head+i].priority
	})
	idx += q.head
	if q.head > 0 {
		if idx == len(q.items) {
			copy(q.items[q.head-1:], q.items[q.head:])
			idx--
		} else {
			copy(q.items[q.head-1:idx], q.items[q.head:idx+1])
		}
		q.head--
	} else {
		q.items = append(q.items, item)
		copy(q.items[idx+1:], q.items[idx:])
	}
	q.items[idx] = item
	q.size++
}

// Peek TODO
func (q *SortedQueue) Peek() interface{} {
	q.Lock()
	defer q.Unlock()

	if len(q.items) == 0 {
		return nil
	}
	item := q.items[q.head]
	return item.value
}

// Pop TODO
func (q *SortedQueue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	if len(q.items) == 0 {
		return nil
	}
	item := q.items[q.head]
	q.items[q.head] = nil
	q.head++
	if q.head > 128 {
		q.items = q.items[q.head:]
		q.head = 0
	}
	q.size--
	return item.value
}

// Size TODO
func (q *SortedQueue) Size() int {
	q.Lock()
	defer q.Unlock()

	return q.size
}

type sortedQueueItem struct {
	value    interface{}
	priority uint64
}

var sortedQueueItemPool = sync.Pool{
	New: func() interface{} {
		return &sortedQueueItem{}
	},
}
