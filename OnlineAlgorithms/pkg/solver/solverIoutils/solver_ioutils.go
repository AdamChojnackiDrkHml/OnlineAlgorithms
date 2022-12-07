package ioutils

import "fmt"

func DebugPrint(s string, p bool) {
	if p {
		fmt.Print(s)
	}
}
