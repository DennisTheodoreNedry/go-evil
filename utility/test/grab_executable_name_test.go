package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestGrabExecutableName(t *testing.T) {
	result := tools.Grab_executable_name()

	assert.Equal(t, filepath.Base(os.Args[0]), result)
}
