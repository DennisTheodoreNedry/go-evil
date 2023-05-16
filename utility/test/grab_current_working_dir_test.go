package test

import (
	"os"
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestGrabPWD(t *testing.T) {
	result := tools.Grab_CWD()
	path, _ := os.Getwd()

	assert.Equal(t, path, result)
}
