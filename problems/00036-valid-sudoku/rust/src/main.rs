#![allow(dead_code)]
struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut row_map = HashMap::new();
        let mut col_map = HashMap::new();
        for a in 0..9 {
            for b in 0..9 {
                match board[a][b] {
                    '.' => {}
                    x => {
                        row_map.entry(x).and_modify(|val| *val += 1).or_insert(1);
                    }
                }
                match board[b][a] {
                    '.' => {}
                    x => {
                        col_map.entry(x).and_modify(|val| *val += 1).or_insert(1);
                    }
                }
            }
            if row_map.values().any(|val| val > &1) || col_map.values().any(|val| val > &1) {
                return false;
            }
            row_map.clear();
            col_map.clear();
        }
        let mut grid_map = HashMap::new();
        for a in (1..9).step_by(3) {
            for b in (1..9).step_by(3) {
                #[allow(clippy::needless_range_loop)]
                for i in (a - 1)..(a + 2) {
                    for j in (b - 1)..(b + 2) {
                        match board[i][j] {
                            '.' => {}
                            x => {
                                grid_map.entry(x).and_modify(|val| *val += 1).or_insert(1);
                            }
                        }
                    }
                }
                if grid_map.values().any(|val| val > &1) {
                    return false;
                }
                grid_map.clear();
            }
        }
        true
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}

