package main

import "math"

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// [3,9,20,null,null,15,7] -> 3
// [1,null,2] -> 2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	return maxDepthHelper(root, 0)
}

func maxDepthHelper(root *TreeNode, currentDepth int) int {
	if root == nil {
		return currentDepth
	}

	currentDepth++
	return int(math.Max(
		float64(maxDepthHelper(root.Left, currentDepth)),
		float64(maxDepthHelper(root.Right, currentDepth)),
	))
}

// Time complexity: O(n). Worst case scenario if we have a single branch.
// Space complexity: O(n). Depth of the recursion stack, worst case scenario if we have a single branch.
