package updatelistsolver

import (
	"OnlineAlgorithms/internal/utils"
	"fmt"
)

type UpdateListAlg int

const (
	MTF UpdateListAlg = iota
	TRANS
	FQ
	BIT
	TS
)

func (e UpdateListAlg) String() string {
	switch e {
	case MTF:
		return "MTF"
	case TRANS:
		return "TRANS"
	case FQ:
		return "FQ"
	case BIT:
		return "BIT"
	case TS:
		return "TS"
	default:
		return "NULL"
	}
}

type UpdateListSolver struct {
	size int
	cost int
	alg  UpdateListSolvingAlg
	algE UpdateListAlg
}

func UpdateListSolver_Create(size int, alg int, debug bool) *UpdateListSolver {
	uLS := &UpdateListSolver{size: size, cost: 0, algE: UpdateListAlg(alg)}
	uLS.createSolvingAlg(alg, debug)
	return uLS
}

func (uLS *UpdateListSolver) Serve(request int) {
	uLS.cost += uLS.alg.UpdateList(request)
}

func (uLS *UpdateListSolver) createSolvingAlg(alg int, debug bool) {
	switch UpdateListAlg(alg) {
	case MTF:
		{
			uLS.alg = MTFAlg_Create(uLS.size, debug)
			break
		}
	case TRANS:
		{
			uLS.alg = TransAlg_Create(uLS.size, debug)
			break
		}
	case FQ:
		{
			uLS.alg = FQAlg_Create(uLS.size, debug)
			break
		}
	case BIT:
		{
			uLS.alg = BITAlg_Create(uLS.size, debug)
			utils.DebugPrint("DUPA BIT\n", debug)
			break
		}
	case TS:
		{
			uLS.alg = TSAlg_Create(uLS.size, debug)
			utils.DebugPrint("DUPA TS\n", debug)

			break
		}
	}
}

func (uLS *UpdateListSolver) Raport() (string, int) {
	return fmt.Sprint(uLS.algE), uLS.cost
}
