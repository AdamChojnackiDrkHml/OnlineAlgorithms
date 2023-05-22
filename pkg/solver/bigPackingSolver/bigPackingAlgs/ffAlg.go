package bigpackingalgs

type FFAlg struct {
	cups  []float64
	debug bool
}

func FFAlg_Create(debug bool) *FFAlg {
	return &FFAlg{cups: make([]float64, 1), debug: debug}
}

func (alg *FFAlg) AddElem(elem float64) {

	for i := range alg.cups {
		if alg.cups[i]+elem <= 1.0 {
			alg.cups[i] += elem
			return
		}
	}

	alg.cups = append(alg.cups, elem)
}

func (alg *FFAlg) GetCups() int {
	return len(alg.cups)
}

func (alg *FFAlg) Clear() {
	alg.cups = make([]float64, 1)
}
