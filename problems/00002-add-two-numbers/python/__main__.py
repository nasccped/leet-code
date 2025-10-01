# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: Optional[ListNode]
        :type l2: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        if not l1 and not l2:
            return None
        l1, l2 = (l1, l2) if l1 else (l2, l1)
        if l2:
            l1.val += l2.val
            l2 = l2.next
        rem = l1.val - (l1.val % 10)
        l1.val -= rem
        rem /= 10
        if l2:
            l2.val += rem
        elif rem > 0:
            l2 = ListNode(rem)

        l1.next = self.addTwoNumbers(l1.next, l2)
        return l1
