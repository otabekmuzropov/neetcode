package main

func main() {

}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}

	return nil
}
