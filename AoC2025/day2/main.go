package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(giftShop())
}

func giftShop() int64 {
	var sum int64

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting wd", err)
	}
	file, err := os.Open(dir + "/AoC2025/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// var positions string

	fileByte, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("error read all", err)
	}

	positions := strings.Split(string(fileByte), ",")
	fmt.Println(len(positions), positions)

	for _, lr := range positions {
		nums := strings.Split(lr, "-")
		right, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			log.Fatal("cannot parse right", nums[1], err)
		}

		left, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			log.Fatal("cannot parse left", nums[0], err)
		}

		for i := left; i <= right; i++ {
			iStr := fmt.Sprintf("%d", i)
			l := len(iStr)
			if l%2 == 1 {
				continue
			}

			if iStr[:l/2] == iStr[l/2:] {
				sum += i
			}
		}
	}

	return sum
}

func Divisors(n int) []int {
	if n <= 0 {
		return []int{}
	}

	divs := []int{}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			if i != n/i {
				divs = append(divs, n/i)
			}
		}
	}
	return divs
}
