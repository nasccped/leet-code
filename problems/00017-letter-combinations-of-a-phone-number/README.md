# 17. Letter Combinations Of A Phone Number

[![Medium](https://img.shields.io/badge/Medium-916f31)](#)
[![Hash Table](https://img.shields.io/badge/Hash_Table-302f33)](#)
[![String](https://img.shields.io/badge/String-302f33)](#)
[![Backtracking](https://img.shields.io/badge/Backtracking-302f33)](#)

Given a string containing digits from `2-9` inclusive, return all
possible letter combinations that the number could represent. Return
the answer in **any order**.

A mapping of digits to letters (just like on the telephone buttons)
is given below. Note that 1 does not map to any letters.

<img src="./example-01.png" alt="example 1" width="350px">

**Example 1:**

> **Input:** digits = "23"
>
> **Output:** ["ad","ae","af","bd","be","bf","cd","ce","cf"]

**Example 2:**

> **Input:** digits = ""
>
> **Output:** []

**Example 3:**

> **Input:** digits = "2"
>
> **Output:** ["a","b","c"]

**Constraints:**

- `0 <= digits.length <= 4`
- `digits[i]` is a digit in the range `['2', '9']`.
