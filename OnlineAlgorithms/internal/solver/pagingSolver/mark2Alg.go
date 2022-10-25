package pagingsolver

import (
	"OnlineAlgorithms/internal/utils"
	"fmt"
	"math"
)

type MARK2Alg struct {
	memory []int
	marks  []bool
	fq     []int
	size   int
	debug  bool
}

func MARK2Alg_Create(size int, debug bool) *MARK2Alg {
	return &MARK2Alg{size: size, memory: make([]int, 0), marks: make([]bool, 0), fq: make([]int, 0), debug: debug}
}

func (alg *MARK2Alg) UpdateMemory(request int) bool {
	utils.DebugPrint((fmt.Sprint("looking for ", request, "\t")), alg.debug)
	isFound := alg.find(request)
	utils.DebugPrint(fmt.Sprint(alg.memory, "\t"), alg.debug)
	utils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)
	utils.DebugPrint(fmt.Sprint(alg.fq, "\t"), alg.debug)

	if !isFound {

		utils.DebugPrint(" ## FAULT ", alg.debug)
		utils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			alg.checkAllMarks()
			evictIndex := alg.findSmallestFqUnmarked()

			if evictIndex == -1 {
				utils.ExitWithError("Unexcpeted")
			}

			utils.DebugPrint(fmt.Sprint(" ## POPPING ", alg.memory[evictIndex], " ## "), alg.debug)
			alg.memory = append(alg.memory[:evictIndex], alg.memory[evictIndex+1:]...)
			alg.marks = append(alg.marks[:evictIndex], alg.marks[evictIndex+1:]...)
			alg.fq = append(alg.fq[:evictIndex], alg.fq[evictIndex+1:]...)

		}
		alg.memory = append([]int{request}, alg.memory...)
		alg.marks = append([]bool{true}, alg.marks...)
		alg.fq = append([]int{1}, alg.fq...)
		utils.DebugPrint(fmt.Sprint(" =>> ", alg.memory, "\t"), alg.debug)
		utils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)
		utils.DebugPrint(fmt.Sprint(alg.fq, "\t"), alg.debug)

	} else {

		utils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory, "\t"), alg.debug)
		utils.DebugPrint(fmt.Sprint(alg.marks, "\t"), alg.debug)
		utils.DebugPrint(fmt.Sprint(alg.fq, "\t"), alg.debug)

	}
	utils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *MARK2Alg) find(request int) bool {
	for i, n := range alg.memory {
		if n == request {
			alg.marks[i] = true
			alg.fq[i] += 1
			return true
		}
	}
	return false
}

func (alg *MARK2Alg) checkAllMarks() {
	for _, n := range alg.marks {
		if !n {
			return
		}
	}
	for i := range alg.marks {
		alg.marks[i] = false
	}
}

func (alg *MARK2Alg) findSmallestFqUnmarked() int {
	minIndex := -1
	minValue := math.MaxInt
	for i, n := range alg.fq {
		if !alg.marks[i] && n < minValue {
			minValue = n
			minIndex = i
		}
	}

	return minIndex
}
