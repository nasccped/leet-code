#![allow(dead_code)]
struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn find_substring(s: String, words: Vec<String>) -> Vec<i32> {
        let mut result = Vec::with_capacity(s.len());
        let syll_len = words[0].len();
        let word_len = words.len() * syll_len;
        if s.len() < word_len {
            return Vec::new();
        }
        let word_count = words.into_iter().fold(HashMap::new(), |mut map, word| {
            map.entry(word).and_modify(|e| *e += 1).or_insert(1);
            map
        });
        's_iter: for i in 0..=(s.len() - word_len) {
            let mut cloned = word_count.clone();
            for j in (i..(i + word_len)).step_by(syll_len) {
                match cloned.get_mut(&s[j..(j + syll_len)]) {
                    Some(v) if *v > 0 => {
                        *v -= 1;
                    }
                    _ => continue 's_iter,
                }
            }
            result.push(i as i32);
        }
        result
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}
