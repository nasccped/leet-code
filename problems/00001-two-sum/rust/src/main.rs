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
    println!(
        "This program doesn't have testing since there's no suitable way of complex testing",
    );
}
