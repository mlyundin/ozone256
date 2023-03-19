package utils

import (
	"sync"
	"time"
)

type item[V any] struct {
	value     V
	timestamp time.Time
}

type TimeQueue[V any] struct {
	queue []item[V]
	lock  sync.Mutex
}

func NewTimeQueue[V any]() *TimeQueue[V] {
	return &TimeQueue[V]{
		queue: make([]item[V], 0)}
}

func (q *TimeQueue[V]) Add(value V) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, item[V]{value: value, timestamp: time.Now()})
}

func (q *TimeQueue[V]) Before(timestamp time.Time) []V {
	q.lock.Lock()
	defer q.lock.Unlock()

	pos := len(q.queue)
	for i, item := range q.queue {
		if item.timestamp.After(timestamp) {
			pos = i
			break
		}
	}

	res := make([]V, 0, pos)
	for _, item := range q.queue[:pos] {
		res = append(res, item.value)
	}
	q.queue = q.queue[pos:]

	return res
}
