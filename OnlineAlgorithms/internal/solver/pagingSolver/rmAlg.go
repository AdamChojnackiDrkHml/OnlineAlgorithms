package pagingsolver

import (
	ioutils "OnlineAlgorithms/internal/utils/ioUtils"
	"fmt"
	"math/rand"
	"time"
)

type RMAlg struct {
	memory []int
	marks  []bool
	size   int
	debug  bool
}

func RMAlg_Create(size int, debug bool) *RMAlg {
	return &RMAlg{size: size, memory: make([]int, 0), marks: make([]bool, 0), debug: debug}
}

func (alg *RMAlg) UpdateMemory(request int) bool {
	ioutils.DebugPrint((fmt.Sprint("looking for ", request, "\t")), alg.debug)
	isFound := alg.find(request)
	ioutils.DebugPrint(fmt.Sprint(alg.memory, "\t"), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)

	if !isFound {

		ioutils.DebugPrint(" ## FAULT ", alg.debug)
		ioutils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			alg.checkAllMarks()
			evictIndex := alg.findItemToPop()

			if evictIndex == -1 {
				ioutils.ExitWithError("Unexcpeted")
			}

			ioutils.DebugPrint(fmt.Sprint(" ## POPPING ", alg.memory[evictIndex], " ## "), alg.debug)
			alg.memory = append(alg.memory[:evictIndex], alg.memory[evictIndex+1:]...)
			alg.marks = append(alg.marks[:evictIndex], alg.marks[evictIndex+1:]...)

		}
		alg.memory = append([]int{request}, alg.memory...)
		alg.marks = append([]bool{true}, alg.marks...)
		ioutils.DebugPrint(fmt.Sprint(" =>> ", alg.memory, "\t"), alg.debug)
		ioutils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)

	} else {

		ioutils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory, "\t"), alg.debug)
		ioutils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)

	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *RMAlg) find(request int) bool {
	for i, n := range alg.memory {
		if n == request {
			alg.marks[i] = true
			return true
		}
	}
	return false
}

func (alg *RMAlg) checkAllMarks() {
	for _, n := range alg.marks {
		if !n {
			return
		}
	}
	for i := range alg.marks {
		alg.marks[i] = false
	}
}

func (alg *RMAlg) findItemToPop() int {

	rand.Seed(time.Now().UTC().UnixNano())
	copyMarksIndx := make([]int, 0)

	for i, n := range alg.marks {
		if !n {
			copyMarksIndx = append(copyMarksIndx, i)
		}
	}

	return copyMarksIndx[rand.Intn(len(copyMarksIndx))]

}
