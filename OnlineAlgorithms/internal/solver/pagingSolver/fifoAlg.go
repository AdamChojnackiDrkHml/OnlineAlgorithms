package pagingsolver

import "fmt"

type FIFOAlg struct {
	memory []int
	size   int
}

func FIFOAlg_Create(size int) *FIFOAlg {
	lru := &FIFOAlg{size: size, memory: make([]int, 0)}

	return lru
}

func (alg *FIFOAlg) UpdateMemory(request int) bool {
	isFound := alg.find(request)
	fmt.Print(alg.memory)
	if !isFound {
		fmt.Print(" ## FAULT ")
		fmt.Print(" HAVE TO INSERT ", request, " ## ")
		if len(alg.memory) >= alg.size {
			fmt.Print(" ## POPPING ", alg.memory[alg.size-1], " ## ")
			alg.memory = alg.memory[:alg.size-1]

		}
		alg.memory = append([]int{request}, alg.memory...)
		fmt.Print(" =>> ", alg.memory)
	} else {
		fmt.Print(" ## FOUND ", request, " REQUEST SERVED ## =>> ", alg.memory)
	}
	fmt.Println()
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
