package pagingsolveralgs

// FIFOAlg hods all information for FIFO algorithm.
type FIFOAlg struct {
	memory []int
	size   int
	debug  bool
}

// FIFOAlg_Create takes size and debug flag and initializes FIFO algorithm for Paging.
func FIFOAlg_Create(size int, debug bool) *FIFOAlg {
	return &FIFOAlg{size: size, memory: make([]int, 0), debug: debug}
}

// UpdateMemory is implementation of PagingSolvingAlg interface for FIFO algorithm.
func (alg *FIFOAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)

	if !isFound {
		if len(alg.memory) >= alg.size {
			alg.memory = alg.memory[1:]

		}
		alg.memory = append(alg.memory, request)
	}
	return isFound
}

func (alg *FIFOAlg) find(request int) bool {
	for _, n := range alg.memory {
		if n == request {
			return true
		}
	}
	return false
}
