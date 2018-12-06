package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var filePath = "./input.txt"

type point struct {
	x     int
	y     int
	valid bool
	count int
}

func main() {
	// Part 1
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var points = map[int]*point{}
	minX, maxX, minY, maxY := 9999999, -1, 9999999, -1
	// Find boundaries
	scanner := bufio.NewScanner(file)
	i := 0
	// Map points, find size of the area
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ", ")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		points[i] = &point{x: x, y: y, valid: false, count: 0}
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
		i++
	}

	// check boundaries.
	// A point won't "spill" out the area if for a boundary
	// there's another point on the opposite axis closer than the wall is
	for i, point := range points {
		top, bottom, left, right := false, false, false, false
		for _, other := range points {
			if !top {
				if other.y >= point.y {
					continue
				}
				if calcDistanceX(*point, *other) <= point.y {
					top = true
				}
			}
			if !left {
				if other.x >= point.x {
					continue
				}
				if calcDistanceY(*point, *other) <= point.x {
					left = true
				}
			}
			if !right {
				if other.x <= point.x {
					continue
				}
				if calcDistanceY(*point, *other) <= maxX-point.x {
					right = true
				}
			}
			if !bottom {
				if other.y <= point.y {
					continue
				}
				if calcDistanceX(*point, *other) <= maxY-point.y {
					bottom = true
				}
			}
		}
		if top && bottom && left && right {
			points[i].valid = true
		}
	}

	nbLocationLessThan10k := 0
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			current := point{x: i, y: j, valid: true}
			minDistance := 9999999

			// Part 1: Find closest point for each coord of the map
			poi := -1
			sumDistance := 0
			for k, point := range points {
				distance := calcTaxicabDistance(current, *point)
				// Min distance is shared, nobody wins
				if distance == minDistance {
					poi = -1
				}
				if distance < minDistance {
					poi = k
					minDistance = distance
				}
				// Part 2: sum of distance to all location
				sumDistance += distance
			}
			if poi > -1 {
				points[poi].count++
			}

			if sumDistance < 10000 {
				nbLocationLessThan10k++
			}
		}
	}
	// part 1
	biggest := 0
	for _, point := range points {
		if point.count > biggest && point.valid {
			biggest = point.count
		}
	}
	fmt.Println("Biggest:", biggest)

	// Part 2
	fmt.Println("Nb locations:", nbLocationLessThan10k)
}

func calcTaxicabDistance(from point, to point) (dist int) {
	dist = calcDistanceY(from, to) + calcDistanceX(from, to)
	return
}

func calcDistanceX(from point, to point) int {
	return abs(from.x - to.x)
}

func calcDistanceY(from point, to point) int {
	return abs(from.y - to.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
