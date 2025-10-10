package main

import "slices"

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
func mergeKLists(lists []*ListNode) *ListNode {
	lists = slices.DeleteFunc(lists, func(p *ListNode) bool {
		return p == nil
	})
	if len(lists) == 0 {
		return nil
	}
	var minInd int = -1
	var minNode *ListNode = nil
	for i, node := range lists {
		if minNode == nil || node.Val < minNode.Val {
			minInd, minNode = i, node
		}
	}
	lists[minInd] = lists[minInd].Next
	minNode.Next = mergeKLists(lists)
	return minNode
}
