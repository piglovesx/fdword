package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindExistedWord(t *testing.T) {
	list := []string{"test", "hello"}
	result := FindExistedWord(list)
	assert.Equal(t, 2, len(result))
}

func Test_FindAllLenWords(t *testing.T) {
	list := []string{"test", "hello"}
	result := FindAllLenWords(list)
	fmt.Println(result)
	assert.Equal(t, 4, len(result))
}

func Test_Permute(t *testing.T) {
	list := Permute([]string{"m", "n", "e", "a", "h", "u"})
	fmt.Println(list)
}

func Test_FindAllLenWords_Random(t *testing.T) {
	result := FindAllLenWords(Permute([]string{"m", "n", "e", "a", "h", "u"}))
	fmt.Println(result)
}
