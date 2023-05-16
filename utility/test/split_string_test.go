package test

import (
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestSplitString(t *testing.T) {
	value := tools.Split_string("Hello")

	assert.Equal(t, []string{"H", "e", "l", "l", "o"}, value)
}
