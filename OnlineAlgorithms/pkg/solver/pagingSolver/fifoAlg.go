package pagingsolver

import (
	"OnlineAlgorithms/pkg/solver/utils"
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
	isFound, position := alg.find(request)
	utils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	utils.DebugPrint(fmt.Sprint(" ## LOOKING FOR ", request, " "), alg.debug)

	if !isFound {
		utils.DebugPrint(" ## FAULT ", alg.debug)
		utils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			utils.DebugPrint(fmt.Sprint(" ## POPPING ", alg.memory[alg.size-1], " ## "), alg.debug)
			alg.memory = alg.memory[:alg.size-1]

		}
		alg.memory = append([]int{request}, alg.memory...)
		utils.DebugPrint(fmt.Sprint(" =>> ", alg.memory), alg.debug)
	} else {
		// alg.memory = append(alg.memory[:position], alg.memory[position+1:]...)
		// alg.memory = append([]int{request}, alg.memory...)
		_ = position
		utils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory), alg.debug)
	}
	utils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *FIFOAlg) find(request int) (bool, int) {
	for i, n := range alg.memory {
		if n == request {
			return true, i
		}
	}
	return false, -1
}
