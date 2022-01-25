package main

import (
	"container/heap"
	"time"
)

type QueueItem struct {
	key string
	item *Data
}

func createQueueItem(key string, item *Data)(queueItem *QueueItem) {
	queueItem = &QueueItem{key: key, item: item}
	return
}

type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].item.expirationTime < pq[j].item.expirationTime
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].item.index = i
	pq[j].item.index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  
	item.item.index = -1 
	*pq = old[0:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Data) {
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) checkExpiry(c *InMap){
	for {
		if (*pq).Len() > 0 && (*pq)[0].item.expirationTime < time.Now().Unix() {
			queueItem := heap.Pop(pq).(*QueueItem)
			c.delete(queueItem.key)
		}
	}
}