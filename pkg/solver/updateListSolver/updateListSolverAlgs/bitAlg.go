package updatelistsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"fmt"
	"math/rand"
	"time"
)

// BITMem holds single memory cell for Bit algorithm.
type BITMemCell struct {
	mem int
	bit bool
}

// BItAlg hods all information for Bit algorithm.
type BITAlg struct {
	memory []*BITMemCell
	size   int
	debug  bool
}

// BITAlg_Create takes size and debug flag and initializes Bit algorithm for Update List.
func BITAlg_Create(size int, debug bool) *BITAlg {
	b := &BITAlg{size: size, debug: debug}

	list := createList(size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, n := range list {
		b.memory = append(b.memory, &BITMemCell{mem: n, bit: r.Int()%2 == 0})
	}

	return b
}

// UpdateList is implementation of UpdateListSolvingAlg interface for Bit algorithm.
func (alg *BITAlg) UpdateList(request int) int {
	ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request, ", "), alg.debug)
	for i, n := range alg.memory {
		if n.mem == request {
			ioutils.DebugPrint(fmt.Sprint(" FOUND ", n.mem, " AT INDEX ", i, " "), alg.debug)

			if !n.bit {
				ioutils.DebugPrint("BIT FLIP TO 1, TRANSPOSING TO BEGINING => ", alg.debug)
				alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
				alg.memory = append([]*BITMemCell{n}, alg.memory...)
			}
			ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
			ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
			n.bit = !n.bit
			return i
		}
	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}

func (alg *BITAlg) unpackMemory() [][]int {
	mem := make([][]int, 0)

	for _, n := range alg.memory {
		bit := 0
		if n.bit {
			bit++
		}
		mem = append(mem, []int{n.mem, bit})
	}

	return mem
}
