package pagemigrationsolver

import (
	"OnlineAlgorithms/pkg/graphs"
	pmalgs "OnlineAlgorithms/pkg/solver/pageMigrationSolver/pageMigrationAlgs"
	"fmt"
)

type PageMigrationSolvingAlg interface {
	Request(request uint8)
	GetCost() uint
	Clear()
}

// PageMigrationSolver struct holds specification of problem and choosen algorithm.
type PageMigrationSolver struct {
	alg  PageMigrationSolvingAlg
	algE pmalgs.PageMigrationAlgs
}

// PageMigrationSolver_Create creates PageMigrationSolver struct for given configuration.
// Returns PageMigrationSolver.
func PageMigrationSolver_Create(algPM pmalgs.PageMigrationAlgs, debug bool, g *graphs.Graph) *PageMigrationSolver {
	pMS := &PageMigrationSolver{algE: pmalgs.PageMigrationAlgs(algPM)}
	pMS.createSolvingAlg(algPM, debug, g)
	return pMS
}

func (pMs *PageMigrationSolver) createSolvingAlg(algPM pmalgs.PageMigrationAlgs, debug bool, g *graphs.Graph) {
	switch pmalgs.PageMigrationAlgs(algPM) {
	case pmalgs.MTM:
		{
			pMs.alg = pmalgs.MTMAlg_Create(debug, g)
			break
		}
	case pmalgs.F:
		{
			pMs.alg = pmalgs.FAlg_Create(debug, g)
			break
		}
	}
}

// Serve is implementation of GenericSolver interface
func (pMS *PageMigrationSolver) Serve(request uint8) {
	pMS.alg.Request(request)
}

// Raport is implementation of GenericSolver interface
func (pMS *PageMigrationSolver) Raport() (string, uint) {
	return fmt.Sprint(pMS.algE), pMS.alg.GetCost()
}

func (pMS *PageMigrationSolver) Clear() {
	pMS.alg.Clear()
}
