# 19. Remove Nth Node From End of List

[![Medium](https://img.shields.io/badge/Medium-916f31)](#)
[![Linked List](https://img.shields.io/badge/Linked_List-302f33)](#)
[![Two Pointers](https://img.shields.io/badge/Two_Pointers-302f33)](#)

Given the `head` of a linked list, remove the
<code>n<sup>th</sup></code> node from the end of the list and return
its head.

**Example 1:**

![example 1](./example-01.jpg)

> **Input:** head = [1,2,3,4,5], n = 2
>
> **Output:** [1,2,3,5]

**Example 2:**

> **Input:** head = [1], n = 1
>
> **Output:** []

**Example 3:**

> **Input:** head = [1,2], n = 1
>
> **Output:** [1]

**Constraints:**

- The number of nodes in the list is `sz`.
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

**Follow up:** Could you do this in one pass?
