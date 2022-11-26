package ioutils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlParseFile(t *testing.T) {
	fmt.Println(os.Getwd())
	conf := ReadYamlForConfig("../../data/configs/generic_structure.yml")

	assert.NotNil(t, conf)
}
