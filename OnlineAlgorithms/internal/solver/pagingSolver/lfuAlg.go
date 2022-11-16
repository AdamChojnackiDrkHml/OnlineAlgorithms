package pagingsolver

import (
	ioutils "OnlineAlgorithms/internal/utils/ioUtils"
	"container/heap"
	"fmt"
)

type LFUMemCell struct {
	reqCnt int
	mem    int
	index  int
}

type LFUAlg struct {
	memory PriorityQueueLFU
	size   int
	debug  bool
}

type PriorityQueueLFU []*LFUMemCell

func (pq PriorityQueueLFU) Len() int {
	return len(pq)
}

func (pq PriorityQueueLFU) Less(i, j int) bool {
	return pq[i].reqCnt < pq[j].reqCnt
}

func (pq PriorityQueueLFU) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueLFU) Push(x any) {
	n := len(*pq)
	item := x.(*LFUMemCell)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueLFU) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func LFUAlg_Create(size int, debug bool) *LFUAlg {
	lfu := &LFUAlg{size: size, memory: make(PriorityQueueLFU, 0), debug: debug}
	heap.Init(&lfu.memory)

	return lfu
}

func (pq *PriorityQueueLFU) update(item *LFUMemCell) {
	item.reqCnt++
	heap.Fix(pq, item.index)
}

func (alg *LFUAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)
	ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" ## LOOKING FOR ", request, " "), alg.debug)
	heap.Init(&alg.memory)
	if !isFound {
		ioutils.DebugPrint(" ## FAULT ", alg.debug)
		ioutils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if alg.memory.Len() >= alg.size {
			x := heap.Pop(&alg.memory).(*LFUMemCell)
			ioutils.DebugPrint(fmt.Sprint(" ## POPPING ", x.mem, " ## "), alg.debug)
		}
		heap.Push(&alg.memory, &LFUMemCell{mem: request, reqCnt: 1})
		ioutils.DebugPrint(fmt.Sprint(" =>> ", alg.unpackMemory()), alg.debug)
	} else {
		ioutils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.unpackMemory()), alg.debug)
	}
	heap.Init(&alg.memory)
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
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
