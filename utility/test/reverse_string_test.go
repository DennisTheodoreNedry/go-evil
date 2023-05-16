package test

import (
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	value := "mephisto"
	tools.Reverse_string(&value)

	assert.Equal(t, "otsihpem", value)
}
