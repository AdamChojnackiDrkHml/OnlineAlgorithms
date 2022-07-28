package updatelistsolver

import (
	"OnlineAlgorithms/internal/utils"
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
	utils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	utils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)
	for i, n := range alg.memory {
		if n == request {
			if i == 0 {
				utils.DebugPrint(fmt.Sprint("FOUND ", n, " AT INDEX 0 "), alg.debug)
				utils.DebugPrint(fmt.Sprintln(), alg.debug)
				return i
			}
			utils.DebugPrint(fmt.Sprint(" FOUND ", n, " AT INDEX ", i, " TRANSPOSING WITH ITEM ON POSITION mem[", i-1, "] = ", alg.memory[i-1], " => "), alg.debug)
			alg.memory[i], alg.memory[i-1] = alg.memory[i-1], alg.memory[i]
			utils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
			utils.DebugPrint(fmt.Sprintln(), alg.debug)
			return i
		}
	}
	utils.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}
