package pagingsolver

type PagingSolvingAlg interface {
	UpdateMemory(request int) bool
}
