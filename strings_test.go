package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {

	// strings.Join makes the boundary conditions to general
	names := []string{"zm", "lee"}
	assert.Equal(t, names[1], strings.Join(names[1:], "."))
}
