package client

import (
	"bufio"
	"os"
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

func FindWord(str string, file *os.File) bool {
	// file, _ := os.Open("words_alpha.txt")
	// defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if str == s {
			return true
		}
		if str < s {
			return false
		}
	}
	return false
}

func FindExistedWord(list []string) []string {
	file, _ := os.Open("../words_alpha.txt")
	defer file.Close()
	var result []string
	var searched map[string]bool = make(map[string]bool)
	for _, v := range list {
		if searched[v] {
			continue
		}
		if FindWord(v, file) {
			result = append(result, v)
		}
		searched[v] = true
		file.Seek(0, 0)
	}
	return result
}

func FindAllLenWords(list []string) []string {
	file, _ := os.Open("../words_alpha.txt")
	defer file.Close()
	var result []string
	searched := make(map[string]bool)
	for _, v := range list {
		if searched[v] {
			continue
		}
		for i := len(v); i >= 3; i-- {
			if searched[v[0:i]] {
				continue
			}
			if FindWord(v[0:i], file) {
				result = append(result, v[0:i])
			}
			searched[v[0:i]] = true
			file.Seek(0, 0)
		}
		searched[v] = true
		file.Seek(0, 0)
	}
	return result
}
