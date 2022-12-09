package pagingsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"fmt"
	"math"
)

// MARKFCAlg hods all information for Mark Frequency Count algorithm.
type MARKFCAlg struct {
	memory []int
	marks  []bool
	fq     []int
	size   int
	debug  bool
}

// MARKFCAlg_Create takes size and debug flag and initializes Mark Frequency Count algorithm for Paging.
func MARKFCAlg_Create(size int, debug bool) *MARKFCAlg {
	return &MARKFCAlg{size: size, memory: make([]int, 0), marks: make([]bool, 0), fq: make([]int, 0), debug: debug}
}

// UpdateMemory is implementation of PagingSolvingAlg interface for Mark Frequency Count algorithm.
func (alg *MARKFCAlg) UpdateMemory(request int) bool {
	ioutils.DebugPrint((fmt.Sprint("looking for ", request, "\t")), alg.debug)
	isFound := alg.find(request)
	ioutils.DebugPrint(fmt.Sprint(alg.memory, "\t"), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(alg.fq, "\t"), alg.debug)

	if !isFound {

		ioutils.DebugPrint(" ## FAULT ", alg.debug)
		ioutils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			alg.checkAllMarks()
			evictIndex := alg.findSmallestFqUnmarked()

			ioutils.DebugPrint(fmt.Sprint(" ## POPPING ", alg.memory[evictIndex], " ## "), alg.debug)
			alg.memory = append(alg.memory[:evictIndex], alg.memory[evictIndex+1:]...)
			alg.marks = append(alg.marks[:evictIndex], alg.marks[evictIndex+1:]...)
			alg.fq = append(alg.fq[:evictIndex], alg.fq[evictIndex+1:]...)

		}
		alg.memory = append([]int{request}, alg.memory...)
		alg.marks = append([]bool{true}, alg.marks...)
		alg.fq = append([]int{1}, alg.fq...)
		ioutils.DebugPrint(fmt.Sprint(" =>> ", alg.memory, "\t"), alg.debug)
		ioutils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)
		ioutils.DebugPrint(fmt.Sprint(alg.fq, "\t"), alg.debug)

	} else {

		ioutils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory, "\t"), alg.debug)
		ioutils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)
		ioutils.DebugPrint(fmt.Sprint(alg.fq, "\t"), alg.debug)

	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *MARKFCAlg) find(request int) bool {
	for i, n := range alg.memory {
		if n == request {
			alg.marks[i] = true
			alg.fq[i] += 1
			return true
		}
	}
	return false
}

func (alg *MARKFCAlg) checkAllMarks() {
	for _, n := range alg.marks {
		if !n {
			return
		}
	}
	for i := range alg.marks {
		alg.marks[i] = false
	}
}

func (alg *MARKFCAlg) findSmallestFqUnmarked() int {
	minIndex := 0
	minValue := math.MaxInt
	for i, n := range alg.fq {
		if !alg.marks[i] && n < minValue {
			minValue = n
			minIndex = i
		}
	}

	return minIndex
}
