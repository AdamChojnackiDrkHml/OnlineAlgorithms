package updatelistsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"fmt"
)

// TSMem holds single memory cell for Timestamp(0) algorithm.
type TSMem struct {
	mem        int
	timestamps []int
}

// TSAlg hods all information for Timestamp(0) algorithm.
type TSAlg struct {
	memory []*TSMem
	size   int
	debug  bool
}

// TSAlg_Create takes size and debug flag and initializes Timestamp(0) algorithm for Update List.
func TSAlg_Create(size int, debug bool) *TSAlg {
	b := &TSAlg{size: size, debug: debug}

	list := createList(size)
	for _, n := range list {
		b.memory = append(b.memory, &TSMem{mem: n, timestamps: make([]int, size)})
	}

	return b
}

// UpdateList is implementation of UpdateListSolvingAlg interface for Timestamp(0) algorithm.
func (alg *TSAlg) UpdateList(request int) int {
	ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)

	for i, n := range alg.memory {
		if n.mem == request {
			ioutils.DebugPrint(fmt.Sprint(" FOUND ", n.mem, " AT INDEX ", i, " LOOKING FOR BEST POSITION -"), alg.debug)
			for j := range n.timestamps {
				n.timestamps[j]++
			}
			for j := 0; j < i; j++ {
				if alg.memory[j].timestamps[request] <= 1 {
					//TODO
					ioutils.DebugPrint(fmt.Sprint(" BEST POSITION AT INDEX ", j, " WITH TIMESTAMP ", alg.memory[j].timestamps[i], " =>"), alg.debug)
					temp := alg.memory[i]
					alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
					alg.memory = append(alg.memory[:j], append([]*TSMem{temp}, alg.memory[j:]...)...)

					break
				}
			}
			for k := range alg.memory {
				alg.memory[k].timestamps[n.mem] = 0
			}

			ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
			ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
			return i
		}
	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}

func (alg *TSAlg) unpackMemory() [][]int {
	mem := make([][]int, 0)

	for i, n := range alg.memory {
		mem = append(mem, make([]int, 0))
		mem[i] = append(mem[i], n.mem)
		mem[i] = append(mem[i], n.timestamps...)
	}

	return mem
}
