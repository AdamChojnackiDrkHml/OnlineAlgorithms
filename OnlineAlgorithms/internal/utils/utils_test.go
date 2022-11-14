package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlParse(t *testing.T) {
	fmt.Println(os.Getwd())
	conf, err := ParseYaml("../../data/configs/generic_structure.yml")

	assert.Equal(t, nil, err)

	assert.NotNil(t, conf)
}
