package pagingsolver

import (
	"OnlineAlgorithms/internal/solver"
	"fmt"
)

type FIFOAlg struct {
	memory []int
	size   int
	debug  bool
}

func FIFOAlg_Create(size int, debug bool) *FIFOAlg {
	return &FIFOAlg{size: size, memory: make([]int, 0), debug: debug}
}

func (alg *FIFOAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)
	solver.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	if !isFound {
		solver.DebugPrint(fmt.Sprint(" ## FAULT "), alg.debug)
		solver.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			solver.DebugPrint(fmt.Sprint(" ## POPPING ", alg.memory[alg.size-1], " ## "), alg.debug)
			alg.memory = alg.memory[:alg.size-1]

		}
		alg.memory = append([]int{request}, alg.memory...)
		solver.DebugPrint(fmt.Sprint(" =>> ", alg.memory), alg.debug)
	} else {
		solver.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory), alg.debug)
	}
	solver.DebugPrint(fmt.Sprintln(), alg.debug)
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
