package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if arrStringLen(args) != 1 {
		return
	}

	input := args[0]

	file, _ := os.Open("standard.txt")
	scanner := bufio.NewScanner(bufio.NewReader(file))

	var resultArr []string

	for scanner.Scan() {
		line := scanner.Text()
		resultArr = append(resultArr, line)
	}

	stringArr := Split(input, "\\n")

	for _, s := range stringArr {
		asciiPrint(s, resultArr)
	}
}

func StringLen(s string) int {
	count := 0
	for range s {
		count++
	}

	return count
}

func Split(str, charset string) []string {
	result := make([]string, countCharset(str, charset)+1)
	lenStr := StringLen((str))
	lenCharset := StringLen((charset))

	j := 0
	last := -1 * lenCharset
	for i, ch := range str {
		if ch == rune(charset[0]) {
			if checkMore(str[i:], charset) {
				result[j] = str[last+lenCharset : i]
				j++
				last = i
			}
		}
	}
	result[j] = str[last+lenCharset : lenStr]

	return result
}

func countCharset(str, charset string) int {
	count := 0
	for i, ch := range str {
		if ch == rune(charset[0]) {
			if checkMore(str[i:], charset) {
				count++
			}
		}
	}

	return count
}

func checkMore(str, charset string) bool {
	for i := 0; i < StringLen(charset); i++ {
		if str[i] != charset[i] {
			return false
		}
	}

	return true
}

func asciiPrint(s string, src []string) {
	for i := 0; i <= 8; i++ {
		for _, ch := range s {
			index := int(rune(ch-32) * 9)
			fmt.Printf("%v", (src[index+i]))
		}
		fmt.Println()
	}
}

func arrStringLen(arr []string) int {
	count := 0
	for range arr {
		count++
	}

	return count
}
