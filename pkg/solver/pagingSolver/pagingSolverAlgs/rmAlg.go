package pagingsolveralgs

import (
	"math/rand"
	"time"
)

// RMAlg hods all information for Random Markup algorithm.
type RMAlg struct {
	memory []int
	marks  []bool
	size   int
	debug  bool
	source *rand.Rand
}

// RMAlg_Create takes size and debug flag and initializes Random Markup algorithm for Paging.
func RMAlg_Create(size int, debug bool) *RMAlg {
	return &RMAlg{
		size:   size,
		memory: make([]int, 0),
		marks:  make([]bool, 0),
		debug:  debug,
		source: rand.New(rand.NewSource(time.Now().UTC().UnixNano()))}
}

// UpdateMemory is implementation of PagingSolvingAlg interface for Random Markup algorithm.
func (alg *RMAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)

	if !isFound {

		if len(alg.memory) >= alg.size {
			alg.checkAllMarks()
			evictIndex := alg.findItemToPop()

			alg.memory = append(alg.memory[:evictIndex], alg.memory[evictIndex+1:]...)
			alg.marks = append(alg.marks[:evictIndex], alg.marks[evictIndex+1:]...)

		}
		alg.memory = append([]int{request}, alg.memory...)
		alg.marks = append([]bool{true}, alg.marks...)

	}
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
	copyMarksIndx := make([]int, len(alg.marks))

	counter := 0
	for i, n := range alg.marks {
		if !n {
			copyMarksIndx[counter] = i
			counter++
		}
	}

	return copyMarksIndx[alg.source.Intn(counter)]

}
