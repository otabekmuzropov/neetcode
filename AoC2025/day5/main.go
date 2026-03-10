package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "/AoC2025/day5/input.txt"
	fmt.Println(CafeteriaPartTwo(path))
}

func Cafeteria(path string) int64 {
	var sum int64

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting wd", err)
	}
	file, err := os.Open(dir + path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var positions []string

	med := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		positions = append(positions, line)
		if line == "" {
			med = len(positions)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	ranges := make([][2]int64, 0)

	for i := 0; i < med-1; i++ {
		pos := positions[i]
		splitted := strings.Split(pos, "-")
		left, err := strconv.ParseInt(splitted[0], 10, 64)
		if err != nil {
			log.Fatal("cannot parse left", err)
		}
		right, err := strconv.ParseInt(splitted[1], 10, 64)
		if err != nil {
			log.Fatal("cannot parse right", err)
		}

		ranges = append(ranges, [2]int64{left, right})

	}

	for i := med; i < len(positions); i++ {
		pos := positions[i]
		target, err := strconv.ParseInt(pos, 10, 64)
		if err != nil {
			log.Fatal("cannot parse right", err)
		}
		fmt.Println("target", target)

		for _, arr := range ranges {
			if target >= arr[0] && target <= arr[1] {
				sum++
				break
			}
		}

	}

	return sum
}

func CafeteriaPartTwo(path string) int64 {
	var sum int64

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting wd", err)
	}
	file, err := os.Open(dir + path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var positions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		positions = append(positions, line)

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	med := 10
	ranges := make([][2]int64, 0)

	for i := 0; i < med-1; i++ {
		pos := positions[i]
		splitted := strings.Split(pos, "-")
		left, err := strconv.ParseInt(splitted[0], 10, 64)
		if err != nil {
			log.Fatal("cannot parse left", err)
		}
		right, err := strconv.ParseInt(splitted[1], 10, 64)
		if err != nil {
			log.Fatal("cannot parse right", err)
		}

		ranges = append(ranges, [2]int64{left, right})

	}

	for i := med; i < len(positions); i++ {
		pos := positions[i]
		target, err := strconv.ParseInt(pos, 10, 64)
		if err != nil {
			log.Fatal("cannot parse right", err)
		}

		for _, arr := range ranges {
			if target >= arr[0] && target <= arr[1] {
				sum++
				break
			}
		}

	}

	return sum
}
