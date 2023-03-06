package main

// https://leetcode.com/problems/contains-duplicate/
// [1,2,3,1] -> true
// [1,2,3,4] -> false
// [1,1,1,3,3,4,3,2,4,2] -> true

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)

	for _, n := range nums {
		_, ok := set[n]
		if ok {
			return true
		}

		set[n] = true
	}

	return false
}

// Time complexity: O(n).
// Space complexity: O(n).
