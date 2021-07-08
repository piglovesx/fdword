package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/piglovesx/fdword/client"
)

const LINE_NUM = 370102
const SPLIT_NUM = 1000
const PAGE_NUM = 371

func main() {
	argsWithoutProg := os.Args[1:]
	word := make([]string, len(argsWithoutProg))
	rands := []string{}
	for _, v := range argsWithoutProg {
		if strings.Contains(v, ":") {
			s := strings.Split(v, ":")
			ind, _ := strconv.Atoi(s[0])
			word[ind] = s[1]
		} else {
			rands = append(rands, v)
		}
	}

	list := client.Permute(rands)
	// fmt.Println(list)
	for i := 0; i < len(list); i++ {
		k := 0
		findstr := []string{}
		chars := strings.Split(list[i], "")
		for j := 0; j < len(word); j++ {
			if word[j] == "" {
				findstr = append(findstr, chars[k])
				k++
			} else {
				findstr = append(findstr, word[j])
			}
		}
		// fmt.Println(strings.Join(findstr, ""))
		if findWord(strings.Join(findstr, "")) {
			fmt.Println(strings.Join(findstr, ""))
		}
	}

}

func findWord(str string) bool {
	file, _ := os.Open("words_alpha.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.TrimSpace(str) == strings.TrimSpace(s) {
			return true
		}
		// if strings.TrimSpace(str) < strings.TrimSpace(s) {
		// 	return false
		// }
	}
	return false
}
