package pagingsolveralgs

// LRUMem holds single memory cell for Least Recently Used algorithm.

// LFUAlg hods all information for Least Recently Used algorithm.
type LRUAlg struct {
	memory []int
	size   int
	debug  bool
}

func LRUAlg_Create(size int, debug bool) *LRUAlg {
	lru := &LRUAlg{size: size, memory: make([]int, 0), debug: debug}

	return lru
}

// LFUAlg_Create takes size and debug flag and initializes Least Recently Used algorithm for Paging.
func (alg *LRUAlg) UpdateMemory(request int) bool {
	isFound, index := alg.find(request)
	if !isFound {
		if len(alg.memory) >= alg.size {
			alg.memory = alg.memory[1:]

		}

		alg.memory = append(alg.memory, request)
	} else if index != alg.size-1 {
		alg.memory = append(alg.memory[:index], alg.memory[index+1:]...)
		alg.memory = append(alg.memory, request)
	}
	return isFound
}

func (alg *LRUAlg) find(request int) (bool, int) {
	for i, n := range alg.memory {
		if n == request {
			return true, i
		}
	}
	return false, -1
}
