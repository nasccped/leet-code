# 42. Trapping Rain Water

[![Hard](https://img.shields.io/badge/Hard-913c31)](#)
[![Array](https://img.shields.io/badge/Array-302f33)](#)
[![Two Pointers](https://img.shields.io/badge/Two_Pointers-302f33)](#)
[![Dynamic Programming](https://img.shields.io/badge/Dynamic_Programming-302f33)](#)
[![Stack](https://img.shields.io/badge/Stack-302f33)](#)
[![Monotonic Stack](https://img.shields.io/badge/Monotonic_Stack-302f33)](#)

Given `n` non-negative integers representing an elevation map where
the width of each bar is `1`, compute how much water it can trap
after raining.

**Example 1:**

![example 1](./example-01.png)

> **Input:** height = [0,1,0,2,1,0,1,3,2,1,2,1]
>
> **Output:** 6
>
> **Explanation:** The above elevation map (black section) is
> represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6
> units of rain water (blue section) are being trapped.

**Example 2:**

> **Input:** height = [4,2,0,3,2,5]
>
> **Output:** 9

**Constraints:**

- `n == height.length`
- <code>1 <= n <= 2 * 10<sup>4</sup></code>
- <code>0 <= height[i] <= 10<sup>5</sup></code>
