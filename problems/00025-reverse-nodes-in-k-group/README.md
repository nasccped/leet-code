# 25. Reverse Nodes in k-Group

[![Hard](https://img.shields.io/badge/Hard-913c31)](#)
[![Linked List](https://img.shields.io/badge/Linked_List-302f33)](#)
[![Recursion](https://img.shields.io/badge/Recursion-302f33)](#)

Given the `head` of a linked list, reverse the nodes of the list `k`
at a time, and return _the modified list_.

`k` is a positive integer and is less than or equal to the length of
the linked list. If the number of nodes is not a multiple of `k` then
left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes
themselves may be changed.

**Example 1:**

![example 1](./example-01.jpg)

> **Input:** head = [1,2,3,4,5], k = 2
>
> **Output:** [2,1,4,3,5]

**Example 2:**

![example 2](./example-02.jpg)

> **Input:** head = [1,2,3,4,5], k = 3
>
> **Output:** [3,2,1,4,5]

**Constraints:**

- The number of nodes in the list is `n`.
- `1 <= k <= n <= 5000`
- `0 <= Node.val <= 1000`

**Follow-up:** Can you solve the problem in `O(1)` extra memory
space?
