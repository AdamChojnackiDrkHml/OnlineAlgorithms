package main

import (
	ioutils "OnlineAlgorithms/internal/testingUtils/ioUtils"
	testCtrl "OnlineAlgorithms/internal/testingUtils/testControler"
	genUtils "OnlineAlgorithms/pkg/utils/generalUtils"
	"os"

	"golang.org/x/exp/slices"
)

func main() {

	if ind := slices.Index(os.Args, "-f"); ind != -1 {
		config := ioutils.ReadYamlForConfig(os.Args[ind+1])
		testCtrl.RunTestWithParametersFromFile(config)
	} else if ind := slices.Index(os.Args, "-p"); ind != -1 {
		testCtrl.RunTestForCmdArguments(ioutils.ParseCmd(os.Args[ind+1:]))
	} else { //hand debug case
		genConf := genUtils.GeneralConfigS{NoOfReq: 500, Iterations: 20, Growth: 500, Repeats: 20}
		solverConf := genUtils.SolverConfigS{ProblemType: 0, Size: 30, AlgP: []genUtils.PagingAlg{1}, Debug: false, DoAll: true}
		generatorConf := genUtils.GeneratorConfigS{DistributionType: []genUtils.GeneratorTypeEnum{0}, Minimum: 0, Maximum: 150, DoAll: false, FvalueGeo: 0.2, FvaluePoiss: 0.3}
		testCtrl.RunTestWithParametersFromFile(&genUtils.Config{TestConfigs: []genUtils.TestConfigS{{GeneralConfig: genConf, SolverConfig: solverConf, GeneratorConfig: generatorConf}}})
	}

}
