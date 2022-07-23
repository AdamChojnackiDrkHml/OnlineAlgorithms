package updatelistsolver

import (
	"OnlineAlgorithms/internal/solver"
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
	solver.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	solver.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)
	for i, n := range alg.memory {
		if n == request {
			if i == 0 {
				solver.DebugPrint(fmt.Sprint("FOUND ", n, " AT INDEX 0 "), alg.debug)
				solver.DebugPrint(fmt.Sprintln(), alg.debug)
				return i
			}
			solver.DebugPrint(fmt.Sprint(" FOUND ", n, " AT INDEX ", i, " TRANSPOSING WITH ITEM ON POSITION mem[", i-1, "] = ", alg.memory[i-1], " => "), alg.debug)
			alg.memory[i], alg.memory[i-1] = alg.memory[i-1], alg.memory[i]
			solver.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
			solver.DebugPrint(fmt.Sprintln(), alg.debug)
			return i
		}
	}
	solver.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}
