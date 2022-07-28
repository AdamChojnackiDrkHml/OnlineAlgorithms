package updatelistsolver

import (
	"OnlineAlgorithms/internal/utils"
	"fmt"
)

type MTFAlg struct {
	memory []int
	size   int
	debug  bool
}

func MTFAlg_Create(size int, debug bool) *MTFAlg {
	return &MTFAlg{size: size, memory: CreateList(size), debug: debug}
}

func (alg *MTFAlg) UpdateList(request int) int {
	utils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	utils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)
	for i, n := range alg.memory {
		if n == request {
			utils.DebugPrint(fmt.Sprint(" FOUND ", n, " AT INDEX ", i, " TRANSPOsING TO BEGINING => "), alg.debug)
			alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
			alg.memory = append([]int{n}, alg.memory...)
			utils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
			utils.DebugPrint(fmt.Sprintln(), alg.debug)
			return i
		}
	}
	utils.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}
