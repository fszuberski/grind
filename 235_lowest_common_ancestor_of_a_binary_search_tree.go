package main

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
// [6,2,8,0,4,7,9,null,null,3,5], 2, 8 -> 6
// [6,2,8,0,4,7,9,null,null,3,5], 2, 4 -> 2
// [2,1], 2, 1 -> 2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

// Time complexity: O(logn). Searching a binary tree
// Space complexity: ~O(n). Depth of recursion stack. Worst case -> one long branch
