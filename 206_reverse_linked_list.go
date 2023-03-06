package main

// https://leetcode.com/problems/reverse-linked-list/description/
// [1,2,3,4,5] -> [5,4,3,2,1]
// [1,2] -> [2, 1]
// [] -> []

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var previous *ListNode = nil
	var next *ListNode
	current := head

	for current != nil {
		next = current.Next
		current.Next = previous

		previous = current
		current = next
	}

	return previous
}

// Time complexity: O(n). Going through each node once.
// Space complexity: O(1). Just keeping pointers to nodes.
