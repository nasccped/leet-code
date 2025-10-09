#![allow(dead_code)]
struct Solution {}

impl Solution {
    pub fn is_valid(s: String) -> bool {
        s.chars()
            .try_fold(Vec::with_capacity(s.len()), |mut stack, character| {
                let repr = match character {
                    '(' | ')' => 'r', // for rounded
                    '[' | ']' => 's', // for squared
                    _ => 'c',         // for curly
                };
                match character {
                    '(' | '[' | '{' => {
                        stack.push(repr);
                        Ok(stack)
                    }
                    _ => {
                        if stack.pop().is_none_or(|ch| ch != repr) {
                            Err(())
                        } else {
                            Ok(stack)
                        }
                    }
                }
            })
            .is_ok_and(|stack| stack.is_empty())
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}
