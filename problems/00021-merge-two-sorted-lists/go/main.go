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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var minNode *ListNode
	if list1.Val < list2.Val {
		minNode = list1
		minNode.Next = mergeTwoLists(list1.Next, list2)
	} else {
		minNode = list2
		minNode.Next = mergeTwoLists(list1, list2.Next)
	}
	return minNode
}
