package maxpriorityqueue

import (
	"introductionToAlgorithms/sort"
	"math"
)

type MaxPriorityQueue interface {
	Insert(item int)
	Max() int
	ExtractMax() int
	IncreaseKey(index int, key int)
	Delete(index int)
	QueueData() []int
}

type maxPriorityqueue struct {
	queue []int
}

func NewMaxPriorityQueue(list []int) MaxPriorityQueue {
	sort.BuildMaxHeapify(list)
	return &maxPriorityqueue{queue: list}
}

func (q *maxPriorityqueue) Insert(item int) {
	q.queue = append(q.queue, math.MinInt32)
	q.IncreaseKey(len(q.queue)-1, item)
}
func (q *maxPriorityqueue) Max() int {
	return q.queue[0]
}

func (q *maxPriorityqueue) ExtractMax() int {
	max := q.queue[0]
	q.queue = q.queue[1:]
	sort.MaxHeapify(q.queue, 0)
	return max
}

func (q *maxPriorityqueue) IncreaseKey(index int, key int) {
	if q.queue[index] < key {
		q.queue[index] = key
		parentIndex := ((index + 1) / 2) - 1
		for index > 0 && q.queue[parentIndex] < q.queue[index] {
			temp := q.queue[parentIndex]
			q.queue[parentIndex] = q.queue[index]
			q.queue[index] = temp
			index = parentIndex
			parentIndex = ((index + 1) / 2) - 1
		}
	}
}

func (q *maxPriorityqueue) Delete(index int) {
	tailItem := q.queue[len(q.queue)-1]
	if q.queue[index] < tailItem {
		q.IncreaseKey(index, tailItem)
	} else {
		q.queue[index] = tailItem
		sort.MaxHeapify(q.queue, index)
	}
	q.queue = q.queue[:len(q.queue)-1]
}

func (q *maxPriorityqueue) QueueData() []int {
	return q.queue
}
