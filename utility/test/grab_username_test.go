package test

import (
	"os/user"
	"testing"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/stretchr/testify/assert"
)

func TestGrabUsername(t *testing.T) {
	result := tools.Grab_username()
	user, _ := user.Current()

	assert.Equal(t, user.Username, result)
}
