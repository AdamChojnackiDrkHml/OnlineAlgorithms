package pagingsolveralgs

// FWFAlg hods all information for FWF algorithm.
type FWFAlg struct {
	memory   []int
	size     int
	debug    bool
	currElem int
}

// FWFAlg_Create takes size and debug flag and initializes FWF algorithm for Paging.
func FWFAlg_Create(size int, debug bool) *FWFAlg {
	return &FWFAlg{size: size,
		memory:   make([]int, size),
		debug:    debug,
		currElem: 0}
}

// UpdateMemory is implementation of PagingSolvingAlg interface for FWF algorithm.
func (alg *FWFAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)

	if !isFound {
		if alg.currElem == alg.size {
			alg.memory = make([]int, alg.size)
			alg.currElem = 0
		}
		alg.memory[alg.currElem] = request
		alg.currElem++
	}
	return isFound
}

func (alg *FWFAlg) find(request int) bool {
	for _, n := range alg.memory {
		if n == request {
			return true
		}
	}
	return false
}
