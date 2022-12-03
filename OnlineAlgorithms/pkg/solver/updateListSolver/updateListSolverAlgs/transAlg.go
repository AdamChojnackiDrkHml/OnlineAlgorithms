package updatelistsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/ioutils"
	"fmt"
)

type TransAlg struct {
	memory []int
	size   int
	debug  bool
}

func TransAlg_Create(size int, debug bool) *TransAlg {
	return &TransAlg{size: size, memory: CreateList(size), debug: debug}
}

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
