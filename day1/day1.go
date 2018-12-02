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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
