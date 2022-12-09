package updatelistsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"fmt"
)

// MTFAlg hods all information for Move-To-Front algorithm.
type MTFAlg struct {
	memory []int
	size   int
	debug  bool
}

// MTFAlg_Create takes size and debug flag and initializes Move-To-Front algorithm for Update List.
func MTFAlg_Create(size int, debug bool) *MTFAlg {
	return &MTFAlg{size: size, memory: createList(size), debug: debug}
}

// UpdateList is implementation of UpdateListSolvingAlg interface for Move-To-Front algorithm.
func (alg *MTFAlg) UpdateList(request int) int {
	ioutils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)
	for i, n := range alg.memory {
		if n == request {
			ioutils.DebugPrint(fmt.Sprint(" FOUND ", n, " AT INDEX ", i, " TRANSPOsING TO BEGINING => "), alg.debug)
			alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
			alg.memory = append([]int{n}, alg.memory...)
			ioutils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
			ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
			return i
		}
	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}
