package common

import "errors"

var errDequeueFromEmptyQueue = errors.New("unable to dequeue from empty queue")
var errPeekFromEmptyQueue = errors.New("unable to peek from empty queue")

type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) Enqueue(el T) {
	q.queue = append(q.queue, el)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.queue) == 0 {
		var defaultValue T
		return defaultValue, errDequeueFromEmptyQueue
	}
	result := q.queue[0]
	q.queue = q.queue[1:]

	return result, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if len(q.queue) == 0 {
		var defaultValue T
		return defaultValue, errPeekFromEmptyQueue
	}
	return q.queue[0], nil
}

func (q *Queue[T]) Len() int {
	return len(q.queue)
}
