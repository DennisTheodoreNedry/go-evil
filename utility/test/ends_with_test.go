package test

import (
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestEndsWith(t *testing.T) {
	result := tools.Ends_with("Hello, world!", []string{"!"})

	ok := result["!"]

	assert.Equal(t, true, ok)
}

func TestDoesntEndWith(t *testing.T) {
	result := tools.Ends_with("Hello, world!", []string{"H"})

	ok := result["H"]

	assert.Equal(t, false, ok)
}
