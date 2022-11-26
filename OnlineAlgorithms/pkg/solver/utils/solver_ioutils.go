package utils

import "fmt"

func DebugPrint(s string, p bool) {
	if p {
		fmt.Print(s)
	}
}
