# 39. Combination Sum

[![Medium](https://img.shields.io/badge/Medium-916f31)](#)
[![Backtracking](https://img.shields.io/badge/Backtracking-302f33)](#)
[![Array](https://img.shields.io/badge/Array-302f33)](#)

Given an array of **distinct** integers `candidates` and a target
integer `target`, return _a list of all **unique combinations** of
`candidates` where the chosen numbers sum to `target`_. You may
return the combinations in **any order**.

The **same** number may be chosen from `candidates` an **unlimited
number of times**. Two combinations are unique if the
frequency[^frequency] of at least one of the chosen numbers is
different.

The test cases are generated such that the number of unique
combinations that sum up to `target` is less than `150` combinations
for the given input.

**Example 1:**

> **Input:** candidates = [2,3,6,7], target = 7
>
> **Output:** [[2,2,3],[7]]
>
> **Explanation:**
>
> 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used
> multiple times.
>
> 7 is a candidate, and 7 = 7.
>
> These are the only two combinations.

**Example 2:**

> **Input:** candidates = [2,3,5], target = 8
>
> **Output:** [[2,2,2,2],[2,3,3],[3,5]]

**Example 3:**

> **Input:** candidates = [2], target = 1
>
> **Output:** []

**Constraints:**

- `1 <= candidates.length <= 30`
- `2 <= candidates[i] <= 40`
- All elements of `candidates` are **distinct**.
- `1 <= target <= 40`

[^frequency]: The frequency of an element `x` is the number of times
  it occurs in the array.
