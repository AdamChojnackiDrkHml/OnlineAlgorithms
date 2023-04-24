package pagingsolveralgs

import (
	"container/heap"
)

// LFUMemCell holds single memory cell for Least Frequently Used algorithm.
type LFUMemCell struct {
	reqCnt int
	mem    int
	index  int
}

// LFUAlg hods all information for Least Frequently Used algorithm.
type LFUAlg struct {
	memory        priorityQueueLFU
	size          int
	debug         bool
	globalCounter map[int]*LFUMemCell
}

// LFUAlg_Create takes size and debug flag and initializes Least Frequently Used algorithm for Paging.
func LFUAlg_Create(size int, debug bool) *LFUAlg {
	lfu := &LFUAlg{
		size:          size,
		memory:        make(priorityQueueLFU, 0),
		debug:         debug,
		globalCounter: make(map[int]*LFUMemCell),
	}
	heap.Init(&lfu.memory)

	return lfu
}

// UpdateMemory is implementation of PagingSolvingAlg interface for Least Frequently Used algorithm.
func (alg *LFUAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)
	heap.Init(&alg.memory)
	if !isFound {
		if alg.memory.Len() >= alg.size {
			x := heap.Pop(&alg.memory).(*LFUMemCell)
			alg.globalCounter[x.mem] = x
		}
		elem, present := alg.globalCounter[request]

		if present {
			//TODO insert
			// heap.Push(&alg.memory, elem)
			elem.reqCnt++
			heap.Push(&alg.memory, elem)
		} else {
			heap.Push(&alg.memory, &LFUMemCell{mem: request, reqCnt: 1})
		}
	}
	heap.Init(&alg.memory)
	return isFound
}

func (pq *priorityQueueLFU) update(item *LFUMemCell) {
	item.reqCnt++
	heap.Fix(pq, item.index)
}

func (alg *LFUAlg) find(request int) bool {
	for _, n := range alg.memory {
		if n.mem == request {
			alg.memory.update(n)
			return true
		}
	}
	return false
}

func (alg *LFUAlg) unpackMemory() [][]int {
	mem := make([][]int, 0)

	for _, n := range alg.memory {
		mem = append(mem, []int{n.mem, n.reqCnt})
	}

	return mem
}

type priorityQueueLFU []*LFUMemCell

func (pq priorityQueueLFU) Len() int {
	return len(pq)
}

func (pq priorityQueueLFU) Less(i, j int) bool {
	return pq[i].reqCnt < pq[j].reqCnt
}

func (pq priorityQueueLFU) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueueLFU) Push(x any) {
	n := len(*pq)
	item := x.(*LFUMemCell)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueueLFU) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
