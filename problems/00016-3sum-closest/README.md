# 16. 3Sum Closest

[![Medium](https://img.shields.io/badge/Medium-916f31)](#)
[![Array](https://img.shields.io/badge/Array-302f33)](#)
[![Two Pointers](https://img.shields.io/badge/Two_Pointers-302f33)](#)
[![Sorting](https://img.shields.io/badge/Sorting-302f33)](#)

Given an integer array `nums` of length `n` and an integer `target`,
find three integers in `nums` such that the sum is closest to
`target`.

Return _the sum of the three integers_.

You may assume that each input would have exactly one solution.

**Example 1:**

> **Input:** nums = [-1,2,1,-4], target = 1
>
> **Output:** 2
>
> **Explanation:** The sum that is closest to the target is 2.
> (-1 + 2 + 1 = 2).

**Example 2:**

> **Input:** nums = [0,0,0], target = 1
>
> **Output:** 0
>
> **Explanation:** The sum that is closest to the target is 0.
> (0 + 0 + 0 = 0).

**Constraints:**

- `3 <= nums.length <= 500`
- `-1000 <= nums[i] <= 1000`
- <code>-10<sup>4</sup> <= target <= 10<sup>4</sup></code>
