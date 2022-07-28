package updatelistsolver

import (
	"OnlineAlgorithms/internal/utils"
	"fmt"
)

type FQAlg struct {
	memory []*FQMem
	size   int
	debug  bool
}

type FQMem struct {
	mem       int
	freqCount int
	index     int
}

func (alg *FQAlg) update(item *FQMem, pos int) {
	item.freqCount++
	if pos == 0 {
		return
	}

	save := pos
	for ; pos > 0 && alg.memory[save].freqCount >= alg.memory[pos-1].freqCount; pos-- {
	}

	alg.memory = append(alg.memory[:save], alg.memory[save+1:]...)
	alg.memory = append(alg.memory[:pos+1], alg.memory[pos:]...)
	alg.memory[pos] = item

}

func FQAlg_Create(size int, debug bool) *FQAlg {
	fq := &FQAlg{size: size, memory: make([]*FQMem, 0), debug: debug}
	list := CreateList(size)

	for i, n := range list {
		fq.memory = append(fq.memory, &FQMem{mem: n, freqCount: 0, index: i})
	}

	return fq
}

func (alg *FQAlg) UpdateList(request int) int {
	utils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
	utils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)
	for i, n := range alg.memory {
		if n.mem == request {

			utils.DebugPrint(fmt.Sprint(" FOUND ", n.mem, " AT INDEX ", i, "UPDATING HEAP => "), alg.debug)
			alg.update(n, i)
			utils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
			utils.DebugPrint(fmt.Sprintln(), alg.debug)
			return i
		}
	}
	utils.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}

func (alg *FQAlg) unpackMemory() [][2]int {
	mem := make([][2]int, 0)

	for _, n := range alg.memory {
		mem = append(mem, [2]int{n.mem, n.freqCount})
	}

	return mem
}
