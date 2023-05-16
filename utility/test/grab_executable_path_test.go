package test

import (
	"os"
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestGrabExecutablePath(t *testing.T) {
	result := tools.Grab_executable_path()

	assert.Equal(t, os.Args[0], result)
}
