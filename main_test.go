package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findWord(t *testing.T) {
	assert.True(t, findWord("test"))
}
