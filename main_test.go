package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findWord(t *testing.T) {
	assert.True(t, findWord("test"))
}

func Test_getLen_num(t *testing.T) {
	os.Args = append(os.Args, "1")

	wordlen, strlen := getLen()
	assert.Equal(t, 4, wordlen)
	assert.Equal(t, 1, strlen)
}

func Test_getLen(t *testing.T) {
	wordlen, strlen := getLen()
	assert.Equal(t, 4, wordlen)
	assert.Equal(t, 4, strlen)
}
