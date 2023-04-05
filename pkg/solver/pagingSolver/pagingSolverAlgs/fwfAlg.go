package pagingsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"fmt"
)

// FWFAlg hods all information for FWF algorithm.
type FWFAlg struct {
	memory []int
	size   int
	debug  bool
}

// FWFAlg_Create takes size and debug flag and initializes FWF algorithm for Paging.
func FWFAlg_Create(size int, debug bool) *FWFAlg {
	return &FWFAlg{size: size, memory: make([]int, 0), debug: debug}
}

// UpdateMemory is implementation of PagingSolvingAlg interface for FWF algorithm.
func (alg *FWFAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)
	ioutils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" ## LOOKING FOR ", request, " "), alg.debug)

	if !isFound {
		ioutils.DebugPrint(" ## FAULT ", alg.debug)
		ioutils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			ioutils.DebugPrint(" ## POPPING ALL ## ", alg.debug)
			alg.memory = make([]int, 0)

		}
		alg.memory = append([]int{request}, alg.memory...)
		ioutils.DebugPrint(fmt.Sprint(" =>> ", alg.memory), alg.debug)
	} else {
		ioutils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory), alg.debug)
	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *FWFAlg) find(request int) bool {
	for _, n := range alg.memory {
		if n == request {
			return true
		}
	}
	return false
}
