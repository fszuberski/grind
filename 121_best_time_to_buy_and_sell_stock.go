package main

import (
	"log"
	"math"
)

func main_121() {
	output := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if output != 5 {
		log.Fatalf("expected: %d", 5)
	}

	output = maxProfit([]int{7, 6, 4, 3, 1})
	if output != 5 {
		log.Fatalf("expected: %d", 0)
	}
}

func maxProfit(prices []int) int {
	// from the back
	// var ans, curMax int
	// for i := len(prices) - 1; i >= 0; i-- {
	//     if curMax - prices[i] > ans {
	//         ans = curMax - prices[i]
	//     }

	//     if prices[i] > curMax {
	//         curMax = prices[i]
	//     }
	// }

	// return ans

	// from the front
	var ans int
	curMin := math.MaxInt
	for i := range prices {
		if prices[i] < curMin {
			curMin = prices[i]
		}

		if prices[i]-curMin > ans {
			ans = prices[i] - curMin
		}
	}

	return ans
}

// Time complexity: O(n). Only a single pass is needed.
// Space complexity: O(1). Only two variables are used.
