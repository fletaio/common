package queue

import (
	"errors"
	"sync"
)

// Queue TODO
type Queue struct {
	sync.Mutex
	pages []*queuePage
	size  int
}

// NewQueue TODO
func NewQueue() *Queue {
	q := &Queue{
		pages: make([]*queuePage, 0, 256),
	}
	return q
}

// Push TODO
func (q *Queue) Push(item interface{}) {
	q.Lock()
	defer q.Unlock()

	var page *queuePage
	if len(q.pages) == 0 {
		page = queuePagePool.Get().(*queuePage)
		q.pages = append(q.pages, page)
	} else {
		page = q.pages[len(q.pages)-1]
		if page.Len() == page.Cap() {
			page = queuePagePool.Get().(*queuePage)
			q.pages = append(q.pages, page)
		}
	}
	page.Push(item)
	q.size++
}

// Peek TODO
func (q *Queue) Peek() interface{} {
	q.Lock()
	defer q.Unlock()

	if len(q.pages) == 0 {
		return nil
	}
	page := q.pages[0]
	item := page.Peek()
	return item
}

// Pop TODO
func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	if len(q.pages) == 0 {
		return nil
	}
	page := q.pages[0]
	item := page.Pop()
	if page.Len() == 0 && len(q.pages) > 1 {
		queuePagePool.Put(page)
		q.pages = q.pages[1:]
	}
	q.size--
	return item
}

// Size TODO
func (q *Queue) Size() int {
	q.Lock()
	defer q.Unlock()

	return q.size
}

var queuePagePool = sync.Pool{
	New: func() interface{} {
		return &queuePage{}
	},
}

var errFullQueue = errors.New("full queue")

type queuePage struct {
	queue [1024]interface{}
	head  int
	tail  int
	size  int
}

// Push TODO
func (page *queuePage) Push(item interface{}) error {
	if page.Len() >= page.Cap() {
		return errFullQueue
	}
	page.queue[page.tail] = item
	page.tail = (page.tail + 1) % page.Cap()
	page.size++
	return nil
}

// Peek TODO
func (page *queuePage) Peek() interface{} {
	if page.Len() == 0 {
		return nil
	}
	item := page.queue[page.head]
	return item
}

// Pop TODO
func (page *queuePage) Pop() interface{} {
	if page.Len() == 0 {
		return nil
	}
	item := page.queue[page.head]
	page.head = (page.head + 1) % page.Cap()
	page.size--
	return item
}

// Len TODO
func (page *queuePage) Len() int {
	return page.size
}

// Cap TODO
func (page *queuePage) Cap() int {
	return len(page.queue)
}
