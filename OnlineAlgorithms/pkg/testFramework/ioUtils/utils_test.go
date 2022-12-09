package testioutils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlParseFile(t *testing.T) {
	fmt.Println(os.Getwd())
	conf := ReadYamlForConfig("../../../data/configs/generic_structure.yaml")

	assert.NotNil(t, conf)
}
