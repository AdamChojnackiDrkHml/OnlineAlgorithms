package updatelistsolver

import (
	"OnlineAlgorithms/internal/utils"
	"fmt"
	"math/rand"
	"time"
)

type BITMem struct {
	mem int
	bit bool
}

type BITAlg struct {
	memory []*BITMem
	size   int
	debug  bool
}

func BITAlg_Create(size int, debug bool) *BITAlg {
	b := &BITAlg{size: size, debug: debug}

	list := CreateList(size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, n := range list {
		b.memory = append(b.memory, &BITMem{mem: n, bit: r.Int()%2 == 0})
	}

	return b
}

func (alg *BITAlg) UpdateList(request int) int {
	utils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
	utils.DebugPrint(fmt.Sprint(" LOOKING FOR ", request, ", "), alg.debug)
	for i, n := range alg.memory {
		if n.mem == request {
			utils.DebugPrint(fmt.Sprint(" FOUND ", n.mem, " AT INDEX ", i, " "), alg.debug)

			if !n.bit {
				utils.DebugPrint("BIT FLIP TO 1, TRANSPOSING TO BEGINING => ", alg.debug)
				alg.memory = append(alg.memory[:i], alg.memory[i+1:]...)
				alg.memory = append([]*BITMem{n}, alg.memory...)
			}
			utils.DebugPrint(fmt.Sprint(alg.unpackMemory()), alg.debug)
			utils.DebugPrint(fmt.Sprintln(), alg.debug)
			n.bit = !n.bit
			return i
		}
	}
	utils.DebugPrint(fmt.Sprintln(), alg.debug)

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
