package test

import (
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	result := tools.Contains("Hello, world!", []string{","})

	ok := result[","]

	assert.Equal(t, true, ok)
}

func TestDoestNotContainValue(t *testing.T) {
	result := tools.Contains("Hello, world!", []string{"."})

	ok := result[","]

	assert.Equal(t, false, ok)
}
