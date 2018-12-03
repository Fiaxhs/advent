package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var filePath = "./input.txt"

const size = 1000

type spot struct {
	overlapCount int
	claims       []int
}

func main() {
	// Part 1
	var canvas = map[int]spot{}
	var nonOverlaping = map[int]bool{}
	overlap := 0
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		id, x, y, w, h := readLine(line)
		nonOverlaping[id] = true
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				where := x + i + size*(y+j)
				currentSpot := canvas[where]
				currentSpot.overlapCount++
				currentSpot.claims = append(currentSpot.claims, id)
				// Count overlap once only (Part1)
				if currentSpot.overlapCount == 2 {
					overlap++
				}
				// Flag overlapped claims (ugly and inefficient)
				if currentSpot.overlapCount > 1 {
					for _, claimID := range currentSpot.claims {
						nonOverlaping[claimID] = false
					}
				}
				canvas[where] = currentSpot
			}
		}
	}
	fmt.Println("Part1:", overlap)
	for id, good := range nonOverlaping {
		if good {
			fmt.Println("Part2:", id)
		}
	}
}

func readLine(line string) (id, x, y, w, h int) {
	// #23 @ 557,102: 28x13
	re := regexp.MustCompile(`#(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+)`)
	matches := re.FindStringSubmatch(line)
	id, _ = strconv.Atoi(matches[1])
	x, _ = strconv.Atoi(matches[2])
	y, _ = strconv.Atoi(matches[3])
	w, _ = strconv.Atoi(matches[4])
	h, _ = strconv.Atoi(matches[5])
	return
}
