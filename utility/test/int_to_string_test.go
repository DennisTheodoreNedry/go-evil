package test

import (
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestIntToString(t *testing.T) {
	result := tools.Int_to_string(666)

	assert.Equal(t, "666", result)
}
