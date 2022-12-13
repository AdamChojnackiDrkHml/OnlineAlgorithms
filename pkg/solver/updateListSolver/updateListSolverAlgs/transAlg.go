package updatelistsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"fmt"
)

// TransAlg hods all information for Transpose algorithm.
type TransAlg struct {
	memory []int
	size   int
	debug  bool
}

// TransAlg_Create takes size and debug flag and initializes Transpose algorithm for Update List.
func TransAlg_Create(size int, debug bool) *TransAlg {
	return &TransAlg{size: size, memory: createList(size), debug: debug}
}

// UpdateList is implementation of UpdateListSolvingAlg interface for Transpose algorithm.
func (alg *TransAlg) UpdateList(request int) int {
	ioutils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)
	for i, n := range alg.memory {
		if n == request {
			if i == 0 {
				ioutils.DebugPrint(fmt.Sprint("FOUND ", n, " AT INDEX 0 "), alg.debug)
				ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
				return i
			}
			ioutils.DebugPrint(fmt.Sprint(" FOUND ", n, " AT INDEX ", i, " TRANSPOSING WITH ITEM ON POSITION mem[", i-1, "] = ", alg.memory[i-1], " => "), alg.debug)
			alg.memory[i], alg.memory[i-1] = alg.memory[i-1], alg.memory[i]
			ioutils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
			ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
			return i
		}
	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}
