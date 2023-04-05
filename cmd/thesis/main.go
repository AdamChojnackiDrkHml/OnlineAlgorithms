package main

import (
	testCtrl "OnlineAlgorithms/pkg/testFramework/testControler"
	"os"
)

func main() {
	testCtrl.RunTestForFileConfig(os.Args[1])
}
