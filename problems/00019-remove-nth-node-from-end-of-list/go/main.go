package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if revIndex(head) == -n {
		return head.Next
	}
	head.Next = removeNthFromEnd(head.Next, n)
	return head
}

// Returns the reverse index (how many nodes to achieve the nil one)
// of the given ListNode pointer.
//
// The returned index will always be negative and start by -1 (to
// mimic Python's reverse indexing: -1, -2, ..., -n). 0 will be
// returned only if the given ListNode pointer is null.
func revIndex(node *ListNode) int {
	if node == nil {
		return 0
	} else {
		return -1 + revIndex(node.Next)
	}
}
