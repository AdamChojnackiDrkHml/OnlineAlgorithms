package solver

import "fmt"

type GenericSolver interface {
	Serve(request int)
	Raport() string
}

func DebugPrint(s string, p bool) {
	if p {
		fmt.Print(s)
	}
}
