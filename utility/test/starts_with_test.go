package test

import (
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestStartsWith(t *testing.T) {
	result := tools.Starts_with("Hello my name is XXXAAAXXX", []string{"Hello"})

	ok := result["Hello"]

	assert.Equal(t, true, ok)
}

func TestDoesntStartWith(t *testing.T) {
	result := tools.Starts_with("Hello, world!", []string{"Peter"})

	ok := result["Peter"]

	assert.Equal(t, false, ok)
}
