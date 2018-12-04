package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

var filePath = "./input.txt"

type day struct {
	guard  int
	events map[int]int
}

type guard struct {
	minutes map[int]int
	total   int
}

func main() {
	// Part 1
	var days = map[string]day{}
	var guards = map[int]guard{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Create logs per date
	for scanner.Scan() {
		line := scanner.Text()
		day, minute, what := readLine(line)
		currentDay := days[day]
		if what == 0 || what == 1 {
			if currentDay.events == nil {
				currentDay.events = map[int]int{minute: what}
			} else {
				currentDay.events[minute] = what
			}
		} else {
			currentDay.guard = what
		}
		days[day] = currentDay
	}

	// Mark and count minutes asleep per guard
	for _, day := range days {
		guard := guards[day.guard]
		awake := true
		for i := 0; i < 60; i++ {
			// if map[i]int == 0 is always true \o/
			if val, ok := day.events[i]; ok && val == 0 {
				awake = false
			} else if day.events[i] == 1 {
				awake = true
			}
			if !awake {
				if guard.minutes == nil {
					guard.minutes = map[int]int{i: 1}
				}
				guard.minutes[i]++
				guard.total++
			}
		}
		guards[day.guard] = guard
	}

	// Find guard with the highest minutes asleep
	max := -1
	sleeperGuard := -1
	for id, guard := range guards {
		if guard.total > max {
			max = guard.total
			sleeperGuard = id
		}
	}
	fmt.Println("Guard sleeping the most", sleeperGuard)
	// Find highest minute slept
	minuteSlept, _ := minuteMostSlept(guards[sleeperGuard].minutes)
	fmt.Println("Minute slept the most", minuteSlept)

	// Part2
	max = -1
	guardFrequently := -1
	minuteFrequently := -1
	for id, guard := range guards {
		minuteSlept, count := minuteMostSlept(guard.minutes)
		if count > max {
			max = count
			guardFrequently = id
			minuteFrequently = minuteSlept
		}
	}
	fmt.Println("Guard most regularily sleeping", guardFrequently)
	fmt.Println("Guard most regularily sleeping, minute slept", minuteFrequently)
}

func readLine(line string) (string, int, int) {
	// [1518-02-23 00:07] falls asleep
	// [1518-09-23 23:53] Guard #2411 begins shift
	// [1518-11-11 00:31] wakes up
	re := regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2}\s\d{2}:\d{2})\].*(wakes|asleep|#(\d+)\s)`)
	// 1 Y-M-D H:i
	// 2 What
	matches := re.FindStringSubmatch(line)
	t, _ := time.Parse("2006-01-02 15:04", matches[1])
	t = t.Add(time.Hour) // Let's not worry about shift starting before midnight
	// fmt.Println(line)
	// fmt.Println(t.Format("0102"))
	// What/id
	// UGLY++++
	what := 0 // sleep
	if matches[2] == "wakes" {
		what = 1
	} else if matches[2] != "sleep" {
		what, _ = strconv.Atoi(matches[3])
	}
	return t.Format("0102"), t.Minute(), what
}

func minuteMostSlept(minutes map[int]int) (int, int) {
	max := 0
	minuteSlept := 0
	for i, count := range minutes {
		if count > max {
			minuteSlept = i
			max = count
		}
	}
	return minuteSlept, max
}
