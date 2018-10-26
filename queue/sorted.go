package queue

import (
	"sort"
	"sync"
)

// SortedQueue sorts items by the priority
type SortedQueue struct {
	sync.Mutex
	items []*sortedQueueItem
	head  int
	size  int
}

// NewSortedQueue returns a SortedQueue
func NewSortedQueue() *SortedQueue {
	q := &SortedQueue{
		items: make([]*sortedQueueItem, 0, 256),
	}
	return q
}

// Insert inserts the item by the priority
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

// Peek fetch the top item without removing it
func (q *SortedQueue) Peek() interface{} {
	q.Lock()
	defer q.Unlock()

	if q.size == 0 {
		return nil
	}
	item := q.items[q.head]
	return item.value
}

// Pop returns a item at the top of the queue
func (q *SortedQueue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	if q.size == 0 {
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

// Size returns the number of items
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
