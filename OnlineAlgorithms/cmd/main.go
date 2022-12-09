package main

import (
	testCtrl "OnlineAlgorithms/pkg/testFramework/testControler"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func main() {

	if ind := slices.Index(os.Args, "-f"); ind != -1 {
		testCtrl.RunTestForFileConfig(os.Args[ind+1])
	} else if ind := slices.Index(os.Args, "-p"); ind != -1 {
		testCtrl.RunTestForCmdArguments(os.Args[ind+1:])
	} else {
		fmt.Println("No test instance provided")
	}

}
