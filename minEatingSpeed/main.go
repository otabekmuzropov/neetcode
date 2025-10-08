package main

import (
	"fmt"
	"math"
)

func main() {
	piles := []int{312884470}
	h := 968709470
	fmt.Println(minEatingSpeedOld(piles, h))
	fmt.Println(minEatingSpeed(piles, h))
}

func minEatingSpeedOld(piles []int, h int) int {
	maxi := piles[0]
	l := len(piles)

	for i := range l {
		maxi = max(maxi, piles[i])
	}

	for m := 1; m <= maxi; m++ {
		count := calculateHours(piles, m)
		// adders := []int{}

		if count <= h {
			// fmt.Println(adders)
			return m
		}
	}

	return maxi
}

func calculateHours(piles []int, m int) int {
	count := 0

	if m == 0 {
		return math.MaxInt
	}

	for i := range piles {
		add := piles[i] / m
		if piles[i]%m != 0 {
			add++
		}

		count += add
	}

	return count
}

func minEatingSpeed(piles []int, h int) int {
	maxi := piles[0]
	l := len(piles)

	count := 0

	for i := range l {
		maxi = max(maxi, piles[i])
	}

	mini := maxi

	low, high := 0, maxi
	for low <= high {
		mid := low + (high-low)/2
		fmt.Println("mid ", mid, low, high, mini)

		count = calculateHours(piles, mid)
		if count > h {
			low = mid + 1
		} else {
			mini = min(mini, mid)
			high = mid - 1
		}

	}

	return mini
}
