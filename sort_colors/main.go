package main

func main() {
	nums := []int{2, 0, 2, 2, 1, 0}
	buckerSort(nums)
}

func sortColors(nums []int) {
	buckerSort(nums)
}	

func buckerSort(nums []int) {
	count := make([]int, 3)
	for _, val := range nums {
		count[val]++
	}

	i := 0
	for idx := 0; idx < len(count); idx++ {
		for j := 0; j < count[idx]; j++ {
			nums[i] = idx
			i++
		}
	}
}
