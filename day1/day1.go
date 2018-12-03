package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var filePath = "./input.txt"

func main() {
	// Part 1
	total := 0
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		total += i
	}
	fmt.Println(total)

	// Part 2
	frequencies := map[int]int{}
	total = 0
	frequencies[total] = 1
	var shouldLoop = true
	for shouldLoop {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			total += i
			frequencies[total]++
			if frequencies[total] > 1 {
				shouldLoop = false
				fmt.Println(total)
				break
			}
		}
	}

}
