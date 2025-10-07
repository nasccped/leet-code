# 11. Container With Most Water

[![Medium](https://img.shields.io/badge/Medium-916f31)](#)
[![Array](https://img.shields.io/badge/Array-302f33)](#)
[![Two Pointers](https://img.shields.io/badge/Two_Pointers-302f33)](#)
[![Greedy](https://img.shields.io/badge/Greedy-302f33)](#)

You are given an integer array `height` of length `n`. There are `n`
vertical lines drawn such that the two endpoints of the
<code>i<sup>th</sup></code> line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such
that the container contains the most water.

Return _the maximum amount of water a container can store_.

**Notice** that you may not slant the container.

**Example 1:**

![example 01](./example-01.jpg)

> **Input:** height = [1,8,6,2,5,4,8,3,7]
>
> **Output:** 49
>
> **Explanation:** The above vertical lines are represented by array
> [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue
> section) the container can contain is 49.

**Example 2:**

> **Input:** height = [1,1]
>
> **Output:** 1

**Constraints:**
- `n == height.length`
- <code>2 <= n <= 10<sup>5</sup></code>
- <code>0 <= height[i] <= 10<sup>4</sup></code>
