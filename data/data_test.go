package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotNil(t *testing.T) {
	users := Populate()

	assert.NotNil(t, users)
}

func TestValidData(t *testing.T) {
	users := Populate()

	assert.Equal(t, users[0].BookZero, 2.5)
}
