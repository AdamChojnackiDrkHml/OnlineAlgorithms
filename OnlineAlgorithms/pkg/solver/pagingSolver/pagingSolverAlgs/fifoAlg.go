package pagingsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"fmt"
)

// FIFOAlg hods all information for FIFO algorithm.
type FIFOAlg struct {
	memory []int
	size   int
	debug  bool
}

// FIFOAlg_Create takes size and debug flag and initializes FIFO algorithm for Paging.
func FIFOAlg_Create(size int, debug bool) *FIFOAlg {
	return &FIFOAlg{size: size, memory: make([]int, 0), debug: debug}
}

// UpdateMemory is implementation of PagingSolvingAlg interface for FIFO algorithm.
func (alg *FIFOAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)
	ioutils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" ## LOOKING FOR ", request, " "), alg.debug)

	if !isFound {
		ioutils.DebugPrint(" ## FAULT ", alg.debug)
		ioutils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			ioutils.DebugPrint(fmt.Sprint(" ## POPPING ", alg.memory[alg.size-1], " ## "), alg.debug)
			alg.memory = alg.memory[:alg.size-1]

		}
		alg.memory = append([]int{request}, alg.memory...)
		ioutils.DebugPrint(fmt.Sprint(" =>> ", alg.memory), alg.debug)
	} else {
		ioutils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory), alg.debug)
	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *FIFOAlg) find(request int) bool {
	for _, n := range alg.memory {
		if n == request {
			return true
		}
	}
	return false
}
