package bigpackingalgs

import (
	"math/rand"
	"time"
)

type RFAlg struct {
	cups   []float64
	source *rand.Rand
	debug  bool
}

func RFAlg_Create(debug bool) *RFAlg {
	return &RFAlg{
		cups:   make([]float64, 1),
		source: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
		debug:  debug,
	}
}

func (alg *RFAlg) AddElem(elem float64) {
	indexes := alg.filterIndex(elem)

	if len(indexes) == 0 {
		alg.cups = append(alg.cups, elem)
		return
	}

	randIndex := rand.Intn(len(indexes))

	alg.cups[indexes[randIndex]] += elem
}

func (alg *RFAlg) filterIndex(elem float64) []int {
	filterIndexes := make([]int, 0)

	for i, cup := range alg.cups {
		if cup+elem <= 1.0 {
			filterIndexes = append(filterIndexes, i)
		}
	}

	return filterIndexes
}

func (alg *RFAlg) GetCups() int {
	return len(alg.cups)
}

func (alg *RFAlg) Clear() {
	alg.cups = make([]float64, 1)
	// alg.source = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}
