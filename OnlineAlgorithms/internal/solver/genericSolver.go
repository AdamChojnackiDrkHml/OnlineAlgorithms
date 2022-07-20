package solver

type GenericSolver interface {
	Serve(request int)
	Raport() string
}
