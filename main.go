package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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

	// for i := 0; i < 8; i++ {
	// 	for _, ch := range s {
	// 		index := int(rune(ch-32) * 9)
	// 		//fmt.Print(index)
	// 		letter, err := rsl("standard.txt", index+i)
	// 		if err != nil {
	// 			fmt.Println("rsl: ", err)
	// 		}
	// 		fmt.Print(string(letter[:len(letter)-1]))

	// 		fmt.Println()
	// 	}
	// }

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
	for i := 0; i < 8; i++ {
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

func rsl(fn string, n int) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("invalid request: line %d", n)
	}
	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()
	bf := bufio.NewReader(f)
	var line string
	for lnum := 0; lnum < n; lnum++ {
		line, err = bf.ReadString('\n')
		if err == io.EOF {
			switch lnum {
			case 0:
				return "", errors.New("no lines in file")
			case 1:
				return "", errors.New("only 1 line")
			default:
				return "", fmt.Errorf("only %d lines", lnum)
			}
		}
		if err != nil {
			return "", err
		}
	}
	if line == "" {
		return "", fmt.Errorf("line %d empty", n)
	}
	return line, nil
}
