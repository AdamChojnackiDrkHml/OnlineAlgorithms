package pagemigrationalgs

import (
	"OnlineAlgorithms/pkg/graphs"
	"math"
)

type MTMAlg struct {
	g       *graphs.Graph
	counter uint
	seq     []uint8
	cost    uint
}

func MTMAlg_Create(debug bool, g *graphs.Graph) *MTMAlg {
	return &MTMAlg{g: g, counter: 0, seq: make([]uint8, D)}
}

func (alg *MTMAlg) Request(req uint8) {
	dist := uint((*alg.g).Request(req))
	alg.cost += dist
	alg.seq[alg.counter] = req
	alg.counter += 1

	if alg.counter == D {
		alg.cost += D * uint((*alg.g).MoveResource(alg.findMin()))
		alg.counter = 0
	}
}

func (alg *MTMAlg) findMin() uint8 {
	var minimum uint
	minimum = math.MaxUint
	minIndex := uint8(64)

	for i := uint8(0); i < 64; i++ {
		sum := uint(0)

		for _, req := range alg.seq {
			sum += uint((*alg.g).Distance(i, req))
		}

		if sum < minimum {
			minimum = sum
			minIndex = i
		}
	}

	return minIndex
}

func (alg *MTMAlg) GetCost() uint {
	return alg.cost
}

func (alg *MTMAlg) Clear() {
	(*alg.g).MoveResource(0)
	alg.counter = 0
	alg.seq = make([]uint8, D)
	alg.cost = 0
}
