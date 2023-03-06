package main

import "math"

// https://leetcode.com/problems/balanced-binary-tree/
// [3,9,20,null,null,15,7] -> true
// [1,2,2,3,3,null,null,4,4] -> false
// [] -> true

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	balanced, _ := isBalancedHelper(root)
	return balanced
}

func isBalancedHelper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	lBalanced, lHeight := isBalancedHelper(root.Left)
	rBalanced, rHeight := isBalancedHelper(root.Right)
	heightDifference := int(math.Abs(float64(lHeight - rHeight)))
	maxHeight := int(math.Max(float64(lHeight), float64(rHeight)))
	return lBalanced && rBalanced && heightDifference <= 1, maxHeight + 1
}

// Time complexity: O(n). Worst case scenario if we have a single branch.
// Space complexity: O(n). Depth of the recursion stack, worst case scenario if we have a single branch.
