package bigpackingalgs

type NFAlg struct {
	current float64
	m       int
	debug   bool
}

func NFAlg_Create(debug bool) *NFAlg {
	return &NFAlg{current: 0.0, m: 1, debug: debug}
}

func (nf *NFAlg) AddElem(elem float64) {
	if nf.current+elem > 1.0 {
		nf.m += 1
		nf.current = elem
		return
	}

	nf.current += elem
}

func (nf *NFAlg) GetCups() int {
	return nf.m
}

func (nf *NFAlg) Clear() {
	nf.current = 0.0
	nf.m = 1
}
