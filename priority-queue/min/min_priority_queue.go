package minpriorityqueue

import (
	"introductionToAlgorithms/sort"
	"math"
)

type MinPriorityQueue interface {
	Insert(item int)
	Min() int
	ExtractMin() int
	DecreaseKey(index int, key int)
	Delete(index int)
	QueueData() []int
}

type minPriorityqueue struct {
	queue []int
}

func NewMinPriorityQueue(list []int) MinPriorityQueue {
	sort.BuildMinHeapify(list)
	return &minPriorityqueue{queue: list}
}

func (p *minPriorityqueue) Insert(item int) {
	p.queue = append(p.queue, math.MaxInt32)
	p.DecreaseKey(len(p.queue)-1, item)
}

func (p *minPriorityqueue) Min() int {
	return p.queue[0]
}

func (p *minPriorityqueue) ExtractMin() int {
	min := p.queue[0]
	p.queue = p.queue[1:]
	sort.MinHeapify(p.queue, 0)
	return min
}

func (p *minPriorityqueue) DecreaseKey(index int, key int) {
	if p.queue[index] > key {
		p.queue[index] = key
		parentIndex := ((index + 1) / 2) - 1
		if index > 0 && p.queue[parentIndex] > p.queue[index] {
			temp := p.queue[parentIndex]
			p.queue[parentIndex] = p.queue[index]
			p.queue[index] = temp
			index = parentIndex
			parentIndex = ((index + 1) / 2) - 1
		}
	}
}

func (p *minPriorityqueue) Delete(index int) {
	tailItem := p.queue[len(p.queue)-1]
	if p.queue[index] > tailItem {
		p.DecreaseKey(index, tailItem)
	} else {
		p.queue[index] = tailItem
		sort.MinHeapify(p.queue, index)
	}
	p.queue = p.queue[:len(p.queue)-1]
}

func (p *minPriorityqueue) QueueData() []int {
	return p.queue
}
