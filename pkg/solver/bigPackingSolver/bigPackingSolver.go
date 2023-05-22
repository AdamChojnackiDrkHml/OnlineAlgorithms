package bigpackingsolver

import (
	bpalgs "OnlineAlgorithms/pkg/solver/bigPackingSolver/bigPackingAlgs"
	"fmt"
)

type BigPackingSolvingAlg interface {
	AddElem(request float64)
	GetCups() int
	Clear()
}

// BigPackingSolver struct holds specification of problem and choosen algorithm.
type BigPackingSolver struct {
	alg  BigPackingSolvingAlg
	algE bpalgs.BigPackingAlg
}

// BigPackingSolver_Create creates BigPackingSolver struct for given configuration.
// Returns BigPackingSolver.
func BigPackingSolver_Create(algBP bpalgs.BigPackingAlg, debug bool) *BigPackingSolver {
	bPS := &BigPackingSolver{algE: bpalgs.BigPackingAlg(algBP)}
	bPS.createSolvingAlg(algBP, debug)
	return bPS
}

func (bps *BigPackingSolver) createSolvingAlg(algBP bpalgs.BigPackingAlg, debug bool) {
	switch bpalgs.BigPackingAlg(algBP) {
	case bpalgs.NF:
		{
			bps.alg = bpalgs.NFAlg_Create(debug)
			break
		}
	case bpalgs.FF:
		{
			bps.alg = bpalgs.FFAlg_Create(debug)
			break
		}
	case bpalgs.RF:
		{
			bps.alg = bpalgs.RFAlg_Create(debug)
			break
		}
	case bpalgs.BF:
		{
			bps.alg = bpalgs.BFAlg_Create(debug)
			break
		}
	case bpalgs.WF:
		{
			bps.alg = bpalgs.WFAlg_Create(debug)
			break
		}
	}
}

// Serve is implementation of GenericSolver interface
func (bPS *BigPackingSolver) Serve(request float64) {
	bPS.alg.AddElem(request)
}

// Raport is implementation of GenericSolver interface
func (bPS *BigPackingSolver) Raport() (string, int) {
	return fmt.Sprint(bPS.algE), bPS.alg.GetCups()
}

func (bPS *BigPackingSolver) Clear() {
	bPS.alg.Clear()
}
