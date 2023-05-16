package test

import (
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestStringToBool(t *testing.T) {
	result := tools.String_to_boolean("true")

	assert.Equal(t, true, result)

	result = tools.String_to_boolean("false")

	assert.Equal(t, false, result)
}
