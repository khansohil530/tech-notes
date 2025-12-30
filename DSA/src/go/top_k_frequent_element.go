package main

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	buckets := make(map[int][]int)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	var result []int
	for i := len(nums); i >= 0; i-- {
		vals, ok := buckets[i]
		if ok {
			result = append(result, vals...)
		}
	}
	return result[:k]
}
