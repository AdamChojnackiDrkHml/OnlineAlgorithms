package updatelistsolver

import (
	"OnlineAlgorithms/internal/solver"
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
	solver.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	solver.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)
	for i, n := range alg.memory {
		if n == request {
			solver.DebugPrint(fmt.Sprint(" FOUND ", n, " AT INDEX ", i, " TRANSPOsING TO BEGINING => "), alg.debug)
			alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
			alg.memory = append([]int{n}, alg.memory...)
			solver.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
			solver.DebugPrint(fmt.Sprintln(), alg.debug)
			return i
		}
	}
	solver.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}
