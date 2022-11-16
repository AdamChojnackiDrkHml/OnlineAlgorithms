package pagingsolver

import (
	ioutils "OnlineAlgorithms/internal/utils/ioUtils"
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
	ioutils.DebugPrint((fmt.Sprint("looking for ", request, "\t")), alg.debug)
	index, isFound := alg.find(request)
	ioutils.DebugPrint(fmt.Sprint(alg.memory, "\t"), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)

	if !isFound {

		ioutils.DebugPrint(" ## FAULT ", alg.debug)
		ioutils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			alg.checkAllMarks()
			for i := len(alg.memory) - 1; i >= 0; i++ {
				if !alg.marks[i] {
					ioutils.DebugPrint(fmt.Sprint(" ## POPPING ", alg.memory[i], " ## "), alg.debug)
					alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
					alg.marks = append(alg.marks[:i], alg.marks[i+1:]...)
					break
				}
			}

		}
		alg.memory = append([]int{request}, alg.memory...)
		alg.marks = append([]bool{true}, alg.marks...)
		ioutils.DebugPrint(fmt.Sprint(" =>> ", alg.memory, "\t"), alg.debug)
		ioutils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)

	} else {
		alg.memory = append(alg.memory[:index], alg.memory[index+1:]...)
		alg.marks = append(alg.marks[:index], alg.marks[index+1:]...)
		alg.memory = append([]int{request}, alg.memory...)
		alg.marks = append([]bool{true}, alg.marks...)
		ioutils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory, "\t"), alg.debug)
		ioutils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)

	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *MARKAlg) find(request int) (int, bool) {
	for i, n := range alg.memory {
		if n == request {
			alg.marks[i] = true
			return i, true
		}
	}
	return -1, false
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
