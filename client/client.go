package client

import (
	"strings"
)

func Permute(str []string) []string {
	list := make([]string, 0, factorial(len(str)))
	return permute(str, len(str), list)
}

func permute(str []string, n int, list []string) []string {
	if n == 1 {
		// fmt.Println(str)
		return append(list, strings.Join(str, ""))
	} else {
		for i := 0; i < n; i++ {
			swap(str, i, n-1)
			list = permute(str, n-1, list)
			swap(str, i, n-1)
		}
	}
	return list
}

func swap(str []string, i, j int) {
	str[i], str[j] = str[j], str[i]
}

func factorial(n int) int {
	result := 1
	for i := n; i >= 1; i-- {
		result *= i
	}
	return result
}
