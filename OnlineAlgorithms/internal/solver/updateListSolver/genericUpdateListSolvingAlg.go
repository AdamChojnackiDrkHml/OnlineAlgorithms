package updatelistsolver

type UpdateListSolvingAlg interface {
	UpdateMemory(request int) bool
}
