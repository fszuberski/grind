package main

// https://leetcode.com/problems/linked-list-cycle/description/
// Note that pos is not passed as a parameter.
// [3,2,0,-4] -> true
// [1,2] -> true
// [1] -> false

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Recursive
/* func hasCycle(head *ListNode) bool {
    set := make(map[*ListNode]bool)
    return hasCycleHelper(head, set)
}

func hasCycleHelper(node *ListNode, set map[*ListNode]bool) bool {
    if node == nil || node.Next == nil {
        return false
    }

    _, ok := set[node]
    if ok {
        return true
    }

    set[node] = true

    return hasCycleHelper(node.Next, set)
} */

// Time complexity: O(n). Going through all nodes once
// Space complexity: O(n). Depth of recursion stack.

// Iterative
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
	}

	return false
}

// Time complexity: O(n). Going through all nodes once
// Space complexity: O(1). Just keeping track of the slow and fast nodes.
