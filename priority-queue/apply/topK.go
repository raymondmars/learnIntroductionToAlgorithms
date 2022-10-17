package apply

import (
	minpriorityqueue "introductionToAlgorithms/priority-queue/min"
)

func topK(list []int, k int) []int {
	kList := []int{}
	for i := 0; i < k; i++ {
		kList = append(kList, 0)
	}
	minQueue := minpriorityqueue.NewMinPriorityQueue(kList)
	for i := 0; i < len(list); i++ {
		if list[i] > minQueue.Min() {
			minQueue.ExtractMin()
			minQueue.Insert(list[i])
		}
	}
	// sort.HeapSort(minQueue.QueueData())
	return minQueue.QueueData()
}
