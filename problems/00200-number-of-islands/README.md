# 200. Number of Islands

[![Medium](https://img.shields.io/badge/Medium-916f31)](#)
[![Array](https://img.shields.io/badge/Array-302f33)](#)
[![Depth-First Search](https://img.shields.io/badge/Depth_First_Search-302f33)](#)
[![Breadth-First Search](https://img.shields.io/badge/Breadth_First_Search-302f33)](#)
[![Union Find](https://img.shields.io/badge/Union_Find-302f33)](#)
[![Matrix](https://img.shields.io/badge/Matrix-302f33)](#)

Given an `m x n` 2D binary grid `grid` which represents a map of
`'1'`s (land) and `'0'`s (water), return _the number of islands_.

An **island** is surrounded by water and is formed by connecting
adjacent lands horizontally or vertically. You may assume all four
edges of the grid are all surrounded by water.

**Example 1:**

> **Input:** grid = [
>
>   ["1","1","1","1","0"],
>
>   ["1","1","0","1","0"],
>
>   ["1","1","0","0","0"],
>
>   ["0","0","0","0","0"]
>
> ]
>
> **Output:** 1

**Example 2:**

> **Input:** grid = [
>
>   ["1","1","0","0","0"],
>
>   ["1","1","0","0","0"],
>
>   ["0","0","1","0","0"],
>
>   ["0","0","0","1","1"]
>
> ]
>
> **Output:** 3

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`.
