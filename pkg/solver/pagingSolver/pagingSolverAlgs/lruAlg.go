package pagingsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"container/heap"
	"fmt"
)

// LRUMem holds single memory cell for Least Recently Used algorithm.
type LRUMem struct {
	mem     int
	lastReq int
	index   int
}

// LFUAlg hods all information for Least Recently Used algorithm.
type LRUAlg struct {
	memory priorityQueue
	size   int
	debug  bool
}

func LRUAlg_Create(size int, debug bool) *LRUAlg {
	lru := &LRUAlg{size: size, memory: make(priorityQueue, 0), debug: debug}
	heap.Init(&lru.memory)

	return lru
}

// LFUAlg_Create takes size and debug flag and initializes Least Recently Used algorithm for Paging.
func (alg *LRUAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)
	ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" ## LOOKING FOR ", request, " "), alg.debug)
	heap.Init(&alg.memory)
	if !isFound {
		ioutils.DebugPrint(" ## FAULT ", alg.debug)
		ioutils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if alg.memory.Len() >= alg.size {
			x := heap.Pop(&alg.memory).(*LRUMem)
			ioutils.DebugPrint(fmt.Sprint(" ## POPPING ", x.mem, " ## "), alg.debug)
		}
		heap.Push(&alg.memory, &LRUMem{mem: request, lastReq: 0})
		ioutils.DebugPrint(fmt.Sprint(" =>> ", alg.unpackMemory()), alg.debug)
	} else {
		ioutils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.unpackMemory()), alg.debug)
	}
	heap.Init(&alg.memory)
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *LRUAlg) find(request int) bool {
	ret := false
	for _, n := range alg.memory {
		if n.mem == request {
			n.lastReq = 0
			ret = true
			continue
		}
		alg.memory.update(n)
	}
	return ret
}

func (alg *LRUAlg) unpackMemory() [][]int {
	mem := make([][]int, 0)

	for _, n := range alg.memory {
		mem = append(mem, []int{n.mem, n.lastReq})
	}

	return mem
}

type priorityQueue []*LRUMem

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].lastReq > pq[j].lastReq
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*LRUMem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) update(item *LRUMem) {
	item.lastReq++
}
