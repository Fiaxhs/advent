package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"unicode/utf8"
)

var filePath = "./input.txt"

// THIS IS LIKE THE UGLIEST, INNEFICIENT WAY TO SOLVE THE CHALLENGE
// DON'T DO IT.

func main() {
	// Part 1
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		origLine := scanner.Text()
		alphabet := "abcdefghijklmnopqrstuvwxyz"

		fmt.Println("Part1:", fullCompute(origLine))

		// Part 2
		min := 9999999
		for _, letter := range alphabet {
			ruleLetter := fmt.Sprintf("(?i)%s", string(letter))
			fmt.Println(ruleLetter)
			re, _ := regexp.Compile(ruleLetter)
			line := re.ReplaceAllString(origLine, "")
			current := fullCompute(line)
			if current < min {
				min = current
			}
		}
		fmt.Println("Part 2:", min)
	}

}

func fullCompute(line string) (size int) {
	rule := "aA|bB|cC|dD|eE|fF|gG|hH|iI|jJ|kK|lL|mM|nN|oO|pP|qQ|rR|sS|tT|uU|vV|wW|xX|yY|zZ|Aa|Bb|Cc|Dd|Ee|Ff|Gg|Hh|Ii|Jj|Kk|Ll|Mm|Nn|Oo|Pp|Qq|Rr|Ss|Tt|Uu|Vv|Ww|Xx|Yy|Zz"
	matching, _ := regexp.MatchString(rule, line)
	for matching {
		re, _ := regexp.Compile(rule)
		line = re.ReplaceAllString(line, "")
		matching, _ = regexp.MatchString(rule, line)
	}
	return utf8.RuneCountInString(line)
}
