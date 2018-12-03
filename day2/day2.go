package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var filePath = "./input.txt"

func main() {
	// Count for part1
	doubleTotal, tripleTotal := 0, 0
	// Dict for part2
	var permutations = map[int]map[string]string{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Part 1 count double and triple
		word := scanner.Text()
		double, triple := countDuosAndTrios(word)
		doubleTotal += double
		tripleTotal += triple

		// Part 2, build possible word missing a letter
		for i := range word {
			wrd := removeLetter(word, i)
			// first word, fill first level of permutations
			if _, ok := permutations[i]; !ok {
				permutations[i] = map[string]string{}
			}
			// wrd already found! We got the winning combination!
			if _, ok := permutations[i][wrd]; ok {
				fmt.Println(wrd)
				break
			}
			permutations[i][wrd] = word
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(doubleTotal * tripleTotal)
}

func countDuosAndTrios(word string) (double, triple int) {
	double, triple = 0, 0
	var dict = map[rune]int{}
	for _, char := range word {
		dict[char]++
	}
	for _, v := range dict {
		if v == 2 {
			double = 1
		}
		if v == 3 {
			triple = 1
		}
	}

	return
}

func removeLetter(word string, pos int) (wrd string) {
	if pos >= len(word) {
		pos = len(word) - 1
	}
	wrd = word[0:pos] + word[pos+1:len(word)]
	return
}
