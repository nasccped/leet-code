# 18. 4Sum

[![Medium](https://img.shields.io/badge/Medium-916f31)](#)
[![Array](https://img.shields.io/badge/Array-302f33)](#)
[![Two Pointers](https://img.shields.io/badge/Two_Pointers-302f33)](#)
[![Sorting](https://img.shields.io/badge/Sorting-302f33)](#)

Given an array `nums` of `n` integers, return _an array of all the
unique quadruplets_ `[nums[a], nums[b], nums[c], nums[d]]` such that:

- `0 <= a, b, c, d < n`
- `a`, `b`, `c`, and `d` are **distinct**.
- `nums[a] + nums[b] + nums[c] + nums[d] == target`

You may return the answer in **any order**.

**Example 1:**

> **Input:** nums = [1,0,-1,0,-2,2], target = 0
>
> **Output:** [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

**Example 2:**

> **Input:** nums = [2,2,2,2,2], target = 8
>
> **Output:** [[2,2,2,2]]

**Constraints:**

- `1 <= nums.length <= 200`
- <code>-10<sup>9</sup> <= nums[i] <= 10<sup>9</sup></code>
- <code>-10<sup>9</sup> <= target <= 10<sup>9</sup></code>
