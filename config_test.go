package behavior3gen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRevise(t *testing.T) {
	c := Config{}
	err := c.Revise()
	assert.NoError(t, err)
	fmt.Println(c)
}
