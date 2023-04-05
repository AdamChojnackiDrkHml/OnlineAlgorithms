package ioutils

import "fmt"

// DebugPrint is utility provided for solvers to debug in runtime.
// Takes string to print and bool value.
// Usually bool value should be specified by Debug field in SolverConfigS.
func DebugPrint(s string, p bool) {
	if p {
		fmt.Print(s)
	}
}
