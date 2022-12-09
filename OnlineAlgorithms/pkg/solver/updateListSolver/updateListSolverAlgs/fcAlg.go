package updatelistsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"fmt"
)

type FCAlg struct {
	memory []*FCMem
	size   int
	debug  bool
}

type FCMem struct {
	mem       int
	freqCount int
	index     int
}

func (alg *FCAlg) update(item *FCMem, pos int) {
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

func FCAlg_Create(size int, debug bool) *FCAlg {
	FC := &FCAlg{size: size, memory: make([]*FCMem, 0), debug: debug}
	list := CreateList(size)

	for i, n := range list {
		FC.memory = append(FC.memory, &FCMem{mem: n, freqCount: 0, index: i})
	}

	return FC
}

func (alg *FCAlg) UpdateList(request int) int {
	ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)
	for i, n := range alg.memory {
		if n.mem == request {

			ioutils.DebugPrint(fmt.Sprint(" FOUND ", n.mem, " AT INDEX ", i, "UPDATING HEAP => "), alg.debug)
			alg.update(n, i)
			ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
			ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
			return i
		}
	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}

func (alg *FCAlg) unpackMemory() [][]int {
	mem := make([][]int, 0)

	for _, n := range alg.memory {
		mem = append(mem, []int{n.mem, n.freqCount})
	}

	return mem
}
