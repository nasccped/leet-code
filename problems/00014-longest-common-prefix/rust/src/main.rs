#![allow(dead_code)]
struct Solution {}

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let mut prefix: String = strs
            .first()
            .expect("`strs` len is always greater than 0")
            .into();
        for s in strs.into_iter().skip(1) {
            while !s.starts_with(&prefix) {
                let _ = prefix.pop();
                if prefix.is_empty() {
                    return prefix;
                }
            }
        }
        prefix
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}
