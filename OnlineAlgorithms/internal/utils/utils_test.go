package utils

import (
	ioutils "OnlineAlgorithms/internal/utils/ioUtils"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlParseFile(t *testing.T) {
	fmt.Println(os.Getwd())
	conf := ioutils.ReadYamlForConfig("../../data/configs/generic_structure.yml")

	assert.NotNil(t, conf)
}
