package pagingsolver

import (
	"container/heap"
	"fmt"
)

type LRUMem struct {
	mem     int
	lastReq int
	index   int
}

type PriorityQueue []*LRUMem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].lastReq > pq[j].lastReq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*LRUMem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *LRUMem) {
	item.lastReq++
}

type LRUAlg struct {
	memory PriorityQueue
	size   int
}

func LRUAlg_Create(size int) *LRUAlg {
	lru := &LRUAlg{size: size, memory: make(PriorityQueue, 0)}
	heap.Init(&lru.memory)

	return lru
}

func (alg *LRUAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)
	fmt.Print(alg.unpackMemory())
	heap.Init(&alg.memory)
	if !isFound {
		fmt.Print(" ## FAULT ")
		fmt.Print(" HAVE TO INSERT ", request, " ## ")
		if alg.memory.Len() >= alg.size {
			x := heap.Pop(&alg.memory).(*LRUMem)
			fmt.Print(" ## POPPING ", x.mem, " ## ")
		}
		heap.Push(&alg.memory, &LRUMem{mem: request, lastReq: 0})
		fmt.Print(" =>> ", alg.unpackMemory())
	} else {
		fmt.Print(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.unpackMemory())
	}
	heap.Init(&alg.memory)
	fmt.Println()
	return isFound
}

func (alg *LRUAlg) find(request int) bool {
	ret := false
	for _, n := range alg.memory {
		if n.mem == request {
			ret = true
			continue
		}
		alg.memory.update(n)
	}
	return ret
}

func (alg *LRUAlg) unpackMemory() [][2]int {
	mem := make([][2]int, 0)

	for _, n := range alg.memory {
		mem = append(mem, [2]int{n.mem, n.lastReq})
	}

	return mem
}
