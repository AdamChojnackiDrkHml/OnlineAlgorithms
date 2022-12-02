package main

import (
	ioutils "OnlineAlgorithms/internal/testingUtils/ioUtils"
	testCtrl "OnlineAlgorithms/internal/testingUtils/testControler"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func main() {

	if ind := slices.Index(os.Args, "-f"); ind != -1 {
		config := ioutils.ReadYamlForConfig(os.Args[ind+1])
		testCtrl.RunTestWithParametersFromFile(config)
	} else if ind := slices.Index(os.Args, "-p"); ind != -1 {
		testCtrl.RunTestForCmdArguments(ioutils.ParseCmd(os.Args[ind+1:]))
	} else {
		fmt.Println("No test instance provided")
	}

}
