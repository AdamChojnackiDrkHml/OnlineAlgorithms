package updatelistsolveralgs

import (
	ioutils "OnlineAlgorithms/pkg/solver/solverIoutils"
	"fmt"
	"math/rand"
	"time"
)

const CHANCE_FOR_TIMESTAMP = 0.2

// CmbMemCell holds single memory cell for Combination algorithm.
type CmbMemCell struct {
	mem        int
	timestamps []int
	bit        bool
}

// CmbAlg hods all information for Combination algorithm.
type CmbAlg struct {
	memory []*CmbMemCell
	size   int
	debug  bool
}

// CombinationAlg_Create takes size and debug flag and initializes Combination algorithm for Update List.
func CombinationAlg_Create(size int, debug bool) *CmbAlg {
	b := &CmbAlg{size: size, debug: debug}

	list := createList(size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, n := range list {
		b.memory = append(b.memory, &CmbMemCell{mem: n, timestamps: make([]int, size), bit: r.Int()%2 == 0})

	}

	return b
}

// UpdateList is implementation of UpdateListSolvingAlg interface for Combination algorithm.
func (alg *CmbAlg) UpdateList(request int) int {
	ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
	ioutils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request), alg.debug)

	for i, n := range alg.memory {
		if n.mem == request {
			ioutils.DebugPrint(fmt.Sprint(" FOUND ", n.mem, " AT INDEX ", i, " LOOKING FOR BEST POSITION -"), alg.debug)

			rand.Seed(time.Now().UTC().UnixNano())
			wildGuess := rand.Float64()

			if wildGuess <= CHANCE_FOR_TIMESTAMP {
				for j := range n.timestamps {
					n.timestamps[j]++
				}
				for j := 0; j < i; j++ {
					if alg.memory[j].timestamps[request] <= 1 {
						//TODO
						ioutils.DebugPrint(fmt.Sprint(" BEST POSITION AT INDEX ", j, " WITH TIMESTAMP ", alg.memory[j].timestamps[i], " =>"), alg.debug)
						temp := alg.memory[i]
						alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
						alg.memory = append(alg.memory[:j], append([]*CmbMemCell{temp}, alg.memory[j:]...)...)

						break
					}
				}
				for k := range alg.memory {
					alg.memory[k].timestamps[n.mem] = 0
				}
			} else {

				if !n.bit {
					ioutils.DebugPrint("BIT FLIP TO 1, TRANSPOSING TO BEGINING => ", alg.debug)

					alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
					alg.memory = append([]*CmbMemCell{n}, alg.memory...)
				}
				n.bit = !n.bit
			}
			ioutils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
			ioutils.DebugPrint(fmt.Sprintln(), alg.debug)

			return i
		}
	}
	ioutils.DebugPrint(fmt.Sprintln(), alg.debug)

	return alg.size
}

func (alg *CmbAlg) unpackMemory() [][]int {
	mem := make([][]int, 0)

	for i, n := range alg.memory {
		mem = append(mem, make([]int, 0))
		mem[i] = append(mem[i], n.mem)
		mem[i] = append(mem[i], n.timestamps...)
	}

	return mem
}
