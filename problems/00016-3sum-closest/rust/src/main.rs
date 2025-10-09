#![allow(dead_code)]
struct Solution {}

impl Solution {
    pub fn three_sum_closest(nums: Vec<i32>, target: i32) -> i32 {
        let mut closest = None;
        let mut sum: i32;
        let mut nums = nums;
        nums.sort_unstable();
        for i in 0..nums.len() - 2 {
            if i > 0 && nums[i] == nums[i - 1] {
                continue;
            }
            let (mut j, mut k) = (i + 1, nums.len() - 1);
            while j < k {
                sum = nums[i] + nums[j] + nums[k];
                match (sum, closest, target) {
                    (s, _, t) if s == t => return s,
                    (s, None, _) => closest = Some(s),
                    (s, Some(c), t) if t.abs_diff(s) < t.abs_diff(c) => closest = Some(s),
                    _ => {}
                }
                if sum < target {
                    while j < k && nums[j] == nums[j + 1] {
                        j += 1;
                    }
                    j += 1;
                } else {
                    while j < k && nums[k] == nums[k - 1] {
                        k -= 1;
                    }
                    k -= 1;
                }
            }
        }
        closest.unwrap()
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}
