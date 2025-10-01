# 3. Longest Substring Without Repeating Characters

[![Medium](https://img.shields.io/badge/Medium-916f31)](#)
[![Hash Table](https://img.shields.io/badge/Hash_Table-302f33)](#)
[![String](https://img.shields.io/badge/String-302f33)](#)
[![Sliding Window](https://img.shields.io/badge/Sliding_Window-302f33)](#)

Given a string `s`, find the length of the **longest
substring**[^substring] without duplicate characters.

**Example 1:**

> **Input:** s = "abcabcbb"
>
> **Output:** 3
>
> **Explanation:** The answer is "abc", with the length of 3.

**Example 2:**

> **Input:** s = "bbbbb"
>
> **Output:** 1
>
> **Explanation:** The answer is "b", with the length of 1.

**Example 3:**

> **Input:** s = "pwwkew"
>
> **Output:** 3
>
> **Explanation:** The answer is "wke", with the length of 3.
>
> Notice that the answer must be a substring, "pwke" is a subsequence
> and not a substring.

**Constraints:**

- `0 <= s.length <= 5 * 10^4`
- `s` consists of English letters, digits, symbols and spaces.

[^substring]: A substring is a contiguous non-empty sequence of
  characters within a string.
