#![allow(dead_code)]
struct Solution {}

impl Solution {
    pub fn longest_valid_parentheses(s: String) -> i32 {
        let mut result = 0;
        let mut stack = Vec::with_capacity(s.len());
        stack.push(-1);
        for (i, c) in s.chars().enumerate().map(|(ind, ch)| (ind as i32, ch)) {
            match c {
                '(' => stack.push(i),
                _ => {
                    stack.pop();
                    if stack.is_empty() {
                        stack.push(i);
                    } else {
                        result = result.max(i - *stack.last().unwrap_or(&i))
                    }
                }
            }
        }
        result
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}
