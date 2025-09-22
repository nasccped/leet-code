#![allow(dead_code)]
use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut result = Vec::with_capacity(2);
        let mut mapping: HashMap<i32, usize> = HashMap::new();
        for (i, n) in nums.into_iter().enumerate() {
            if let Some(val) = mapping.get(&(target - n)) {
                result.push(i as i32);
                result.push(*val as i32);
                break;
            } else {
                mapping.insert(n, i);
            }
        }
        result
    }
}

fn main() {
    let turn_green = |x: &str| format!("\x1b[92m{x}\x1b[0m");
    println!(
        "This program should be ran using `{}`.",
        turn_green("cargo test")
    );
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_01() {
        let (input, target) = ([2, 7, 11, 15], 9);
        let output = [0, 1];
        assert_eq!(
            {
                let mut s = Solution::two_sum(input.into_iter().collect(), target);
                s.sort();
                s
            },
            output
        );
    }

    #[test]
    fn test_02() {
        let (input, target) = ([3, 2, 4], 6);
        let output = [1, 2];
        assert_eq!(
            {
                let mut s = Solution::two_sum(input.into_iter().collect(), target);
                s.sort();
                s
            },
            output
        );
    }

    #[test]
    fn test_03() {
        let (input, target) = ([3, 3], 6);
        let output = [0, 1];
        assert_eq!(
            {
                let mut s = Solution::two_sum(input.into_iter().collect(), target);
                s.sort();
                s
            },
            output
        );
    }
}
