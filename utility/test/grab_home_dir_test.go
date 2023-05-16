package test

import (
	"os/user"
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestGrabHomeDir(t *testing.T) {
	result := tools.Grab_home_dir()
	path, _ := user.Current()

	assert.Equal(t, path.HomeDir, result)
}
