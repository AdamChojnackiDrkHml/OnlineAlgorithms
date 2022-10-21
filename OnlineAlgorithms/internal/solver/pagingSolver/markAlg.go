package pagingsolver

import (
	"OnlineAlgorithms/internal/utils"
	"fmt"
)

type MARKAlg struct {
	memory []int
	marks  []bool
	size   int
	debug  bool
}

func MARKAlg_Create(size int, debug bool) *MARKAlg {
	return &MARKAlg{size: size, memory: make([]int, 0), marks: make([]bool, 0), debug: debug}
}

func (alg *MARKAlg) UpdateMemory(request int) bool {
	utils.DebugPrint((fmt.Sprint("looking for ", request, "\t")), alg.debug)
	isFound := alg.find(request)
	utils.DebugPrint(fmt.Sprint(alg.memory, "\t"), alg.debug)
	utils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)

	alg.checkAllMarks()
	if !isFound {
		utils.DebugPrint(" ## FAULT ", alg.debug)
		utils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			for i, n := range alg.marks {
				if !n {
					utils.DebugPrint(fmt.Sprint(" ## POPPING ", alg.memory[i], " ## "), alg.debug)
					alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
					alg.marks = append(alg.marks[:i], alg.marks[i+1:]...)
					break
				}
			}

		}
		alg.memory = append([]int{request}, alg.memory...)
		alg.marks = append([]bool{true}, alg.marks...)
		utils.DebugPrint(fmt.Sprint(" =>> ", alg.memory, "\t"), alg.debug)
		utils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)

	} else {
		utils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory, "\t"), alg.debug)
		utils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)

	}
	utils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *MARKAlg) find(request int) bool {
	for i, n := range alg.memory {
		if n == request {
			alg.marks[i] = true
			return true
		}
	}
	return false
}

func (alg *MARKAlg) checkAllMarks() {
	for _, n := range alg.marks {
		if !n {
			return
		}
	}
	for i := range alg.marks {
		alg.marks[i] = false
	}
}
