package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	name := "/AoC2025/day1/input.txt"
	fmt.Println(secretEntrance(name))
	fmt.Println(secretEntrance2(name))
}

func secretEntrance(name string) int {
	sum := 0

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting wd", err)
	}
	file, err := os.Open(dir + name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var positions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		positions = append(positions, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	pointer := 50

	for _, pos := range positions {
		direction := pos[0]
		// point := pos[1:]
		point, _ := strconv.Atoi(pos[1:])

		switch direction {
		case 'R':
			pointer = (pointer + point) % 100
		case 'L':
			pointer = ((pointer-point)%100 + 100) % 100
		}

		if pointer == 0 {
			sum++
		}
	}

	return sum
}

func secretEntrance2(name string) int {
	sum := 0

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting wd", err)
	}
	file, err := os.Open(dir + name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var positions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		positions = append(positions, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	pointer := 50

	for _, pos := range positions {
		direction := pos[0]
		// point := pos[1:]
		point, _ := strconv.Atoi(pos[1:])

		switch direction {
		case 'R':
			pointer += point
			sum += pointer / 100
			pointer %= 100
		case 'L':
			pointer = pointer - point

			// pointer = ((pointer-point)%100 + 100) % 100
			sum = sum - pointer/100
			if pointer < 0 {
				sum++
			}

			pointer = (pointer%100 + 100) % 100
		}
		fmt.Println("DIRECTION: ", string(direction), "point: ", point, "SUM: ", sum, "POINTER: ", pointer)

		// if pointer == 0 {
		// 	sum++
		// }
	}

	return sum
}
