package test

import (
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestStringToInt(t *testing.T) {
	result := tools.String_to_int("666")

	assert.Equal(t, 666, result)
}

func TestStringToIntInvalid(t *testing.T) {
	result := tools.String_to_int("true")

	assert.Equal(t, -1, result)
}
