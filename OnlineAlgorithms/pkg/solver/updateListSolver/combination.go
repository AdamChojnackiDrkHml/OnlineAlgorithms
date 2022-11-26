package updatelistsolver

import (
	ioutils "OnlineAlgorithms/pkg/solver/utils"
	"fmt"
	"math/rand"
	"time"
)

const CHANCE_FOR_TIMESTAMP = 0.2

type CombinationMem struct {
	mem        int
	timestamps []int
	bit        bool
}

type CombinationAlg struct {
	memory []*CombinationMem
	size   int
	debug  bool
}

func CombinationAlg_Create(size int, debug bool) *CombinationAlg {
	b := &CombinationAlg{size: size, debug: debug}

	list := CreateList(size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, n := range list {
		b.memory = append(b.memory, &CombinationMem{mem: n, timestamps: make([]int, size), bit: r.Int()%2 == 0})

	}

	return b
}

func (alg *CombinationAlg) UpdateList(request int) int {
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
						alg.memory = append(alg.memory[:j], append([]*CombinationMem{temp}, alg.memory[j:]...)...)

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
					alg.memory = append([]*CombinationMem{n}, alg.memory...)
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

func (alg *CombinationAlg) unpackMemory() [][]int {
	mem := make([][]int, 0)

	for i, n := range alg.memory {
		mem = append(mem, make([]int, 0))
		mem[i] = append(mem[i], n.mem)
		mem[i] = append(mem[i], n.timestamps...)
	}

	return mem
}
