package pagemigrationalgs

import (
	"OnlineAlgorithms/pkg/graphs"
	"math/rand"
	"time"
)

type FAlg struct {
	g    *graphs.Graph
	cost uint
	rand *rand.Rand
}

func FAlg_Create(debug bool, g *graphs.Graph) *FAlg {
	return &FAlg{g: g,
		rand: rand.New(rand.NewSource(time.Now().UTC().UnixNano()))}
}

func (alg *FAlg) Request(req uint8) {
	dist := (*alg.g).Request(req)

	alg.cost += uint(dist)

	if alg.rand.Float64() < 1.0/float64(D) {
		alg.cost += D * uint((*alg.g).MoveResource(req))
	}
}

func (alg *FAlg) GetCost() uint {
	return alg.cost
}

func (alg *FAlg) Clear() {
	(*alg.g).MoveResource(0)
	alg.cost = 0
}
