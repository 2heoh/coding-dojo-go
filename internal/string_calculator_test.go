package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringCalculator_EmptyString(t *testing.T) {
	calculator := StringCalculator{}

	sum := calculator.Add("")
	assert.Equal(t, 0, sum)
}
