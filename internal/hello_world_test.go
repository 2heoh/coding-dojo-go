package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HelloWorld(t *testing.T) {
	helloWorld := HelloWorld{}

	greeting, err := helloWorld.HelloWorld()
	assert.NoError(t, err)

	assert.Equal(t, "Hello, World!", greeting)
}
