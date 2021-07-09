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
	wordlen, strlen := getLen()
	word := make([]string, wordlen)
	rands := []string{}
	for i, v := range os.Args[1:] {
		if i == len(word) {
			break
		}
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
	var found = make(map[string]int)
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
		s := strings.Join(findstr[0:strlen], "")
		// fmt.Println(strlen)
		if found[s] == 1 {
			continue
		}
		found[s] = 1
		// fmt.Println(found)
		// fmt.Println(s)
		if findWord(s) {
			fmt.Println(s)
		}
	}

}

func getLen() (int, int) {
	argsWithoutProg := os.Args[1:]
	argsLen := len(argsWithoutProg)
	if v, is := checkNum(argsWithoutProg[argsLen-1]); is {
		return argsLen - 1, v
	}

	return argsLen, argsLen
}

func checkNum(s string) (int, bool) {
	if v, err := strconv.Atoi(s); err == nil {
		return v, true
	}
	return 0, false
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
