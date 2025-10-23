#![allow(dead_code)]
struct Solution {}

impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut filtered = nums
            .iter()
            .enumerate()
            .filter_map(|(i, n)| if n == &target { Some(i) } else { None });
        let (low, high) = match (
            filtered.next().map(|x| x as i32),
            filtered.next_back().map(|x| x as i32),
        ) {
            (Some(l), Some(h)) => (l, h),
            (Some(a), None) | (None, Some(a)) => (a, a),
            _ => (-1, -1),
        };
        Vec::from([low, high])
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}

