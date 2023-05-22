package bigpackingalgs

type WFAlg struct {
	cups  []float64
	debug bool
}

func WFAlg_Create(debug bool) *WFAlg {
	return &WFAlg{cups: make([]float64, 1), debug: debug}
}

func (alg *WFAlg) AddElem(elem float64) {

	choosenCup := 0
	isChoosen := false
	for i := range alg.cups {
		if alg.cups[i]+elem <= 1.0 {
			alg.cups[i] += elem
			choosenCup = i
			isChoosen = true
			break
		}
	}

	if !isChoosen {
		alg.cups = append(alg.cups, elem)
		choosenCup = len(alg.cups) - 1
	}

	for {
		if choosenCup == 0 {
			break
		}

		if alg.cups[choosenCup-1] <= alg.cups[choosenCup] {
			break
		}

		alg.cups[choosenCup-1], alg.cups[choosenCup] = alg.cups[choosenCup], alg.cups[choosenCup-1]

	}
}

func (alg *WFAlg) GetCups() int {
	return len(alg.cups)
}

func (alg *WFAlg) Clear() {
	alg.cups = make([]float64, 1)
}
