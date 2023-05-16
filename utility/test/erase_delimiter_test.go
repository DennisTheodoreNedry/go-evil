package test

import (
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestEraseDelimiter(t *testing.T) {
	result := tools.Erase_delimiter("Hello, world!", []string{","}, -1)

	assert.Equal(t, "Hello world!", result)
}
