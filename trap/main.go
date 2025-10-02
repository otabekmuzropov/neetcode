package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(height))
	fmt.Println(trap2(height))
}

func trap(height []int) int {
	t := len(height)
	if t == 0 {
		return 0
	}

	prefix := make([]int, t)
	suffix := make([]int, t)

	l, r := 0, t-1
	pMax, sMax := 0, 0

	for range t {
		prefix[l] = pMax
		suffix[r] = sMax

		if height[l] > pMax {
			pMax = height[l]
		}

		if height[r] > sMax {
			sMax = height[r]
		}

		l++
		r--
	}

	sum := 0
	for i := 1; i < t-1; i++ {
		m := max(min(prefix[i], suffix[i])-height[i], 0)
		sum += m
	}

	return sum
}

func trap2(height []int) int {
	t := len(height)
	if t <= 1 {
		return 0
	}

	sum := 0

	l, r := 0, t-1
	lMax := height[0]
	rMax := height[t-1]

	for l < r {
		if lMax < rMax {
			l++
			lMax = max(lMax, height[l])
			sum += lMax - height[l]
		} else {
			r--
			rMax = max(rMax, height[r])
			sum += rMax - height[r]

		}
	}

	return sum
}
