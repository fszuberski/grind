package main

// https://leetcode.com/problems/invert-binary-tree/
// [4,2,7,1,3,6,9] -> [4,7,2,9,6,3,1]
// [2,1,3] -> [2,3,1]
// [] -> []

// Recursive
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left != nil {
		root.Left = invertTree(root.Left)
	}

	if root.Right != nil {
		root.Right = invertTree(root.Right)
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	return root
}

// Because of recursion, O(h) function calls will be placed on the stack in the worst case, where hhh is the height of the tree. Because h∈O(n), the space complexity is O(n).

// This can be done iteratively in a BFS manner to reduce space complexity.
