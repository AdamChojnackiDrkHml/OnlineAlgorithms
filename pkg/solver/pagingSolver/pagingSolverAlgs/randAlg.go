package pagingsolveralgs

import (
	"math/rand"
	"time"
)

// RANDAlg hods all information for RAND algorithm.
type RANDAlg struct {
	memory []int
	size   int
	debug  bool
	source *rand.Rand
}

// RANDAlg_Create takes size and debug flag and initializes RAND algorithm for Paging.
func RANDAlg_Create(size int, debug bool) *RANDAlg {
	return &RANDAlg{size: size,
		memory: make([]int, 0),
		debug:  debug,
		source: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}

// UpdateMemory is implementation of PagingSolvingAlg interface for RAND algorithm.
func (alg *RANDAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)
	// ioutils.DebugPrint(fmt.Sprint(alg.memory), alg.debug)
	// ioutils.DebugPrint(fmt.Sprint(" ## LOOKING FOR ", request, " "), alg.debug)

	if !isFound {
		// ioutils.DebugPrint(" ## FAULT ", alg.debug)
		// ioutils.DebugPrint(fmt.Sprint(" HAVE TO INSERT ", request, " ## "), alg.debug)
		if len(alg.memory) >= alg.size {
			deleteIndex := alg.randomIndex()
			// ioutils.DebugPrint(fmt.Sprint(" ## POPPING ", alg.memory[deleteIndex], " ## "), alg.debug)
			alg.memory = append(alg.memory[:deleteIndex], alg.memory[deleteIndex+1:]...)

		}
		alg.memory = append([]int{request}, alg.memory...)
		// ioutils.DebugPrint(fmt.Sprint(" =>> ", alg.memory), alg.debug)
	} else {
		// ioutils.DebugPrint(fmt.Sprint(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory), alg.debug)
	}
	// ioutils.DebugPrint(fmt.Sprintln(), alg.debug)
	return isFound
}

func (alg *RANDAlg) find(request int) bool {
	for _, n := range alg.memory {
		if n == request {
			return true
		}
	}
	return false
}

func (alg *RANDAlg) randomIndex() int {

	return alg.source.Intn(alg.size)
}
