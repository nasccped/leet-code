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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if temp := reverseIndex(head); temp == 0 || temp > -k {
		return head
	}
	var mainHead *ListNode = &ListNode{0, nil}
	nextList := head
	for i := 0; i < k; i++ {
		nextList = nextList.Next
	}
	currentGroup := make([]*ListNode, k)
	aux := head
	for i := range k {
		currentGroup[k-i-1] = aux
		aux = aux.Next
		currentGroup[k-i-1].Next = nil
	}
	aux = mainHead
	for i := range k {
		aux.Next = currentGroup[i]
		aux = aux.Next
	}
	aux.Next = reverseKGroup(nextList, k)
	return mainHead.Next
}

// Returns the reverse index of a ListNode pointer (similar to
// Python's reverse indexing [-1, -2, ...]).
func reverseIndex(node *ListNode) int {
	if node == nil {
		return 0
	} else {
		return -1 + reverseIndex(node.Next)
	}
}
