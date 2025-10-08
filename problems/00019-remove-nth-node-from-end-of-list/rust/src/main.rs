#![allow(dead_code)]

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

struct Solution {}

// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
//
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        match head {
            Some(h) if Self::reverse_index(&head) == -n => h.next,
            Some(mut h) => {
                h.next = Self::remove_nth_from_end(h.next, n);
                Some(h)
            }
            _ => None,
        }
    }

    /// Returns the reverse index (how many nodes to achieve the [`None`] variant)
    /// of the given ListNode pointer.
    ///
    /// The returned index will always be negative and start by -1 (to
    /// mimic Python's reverse indexing: -1, -2, ..., -n). 0 will be
    /// returned only if the given ListNode pointer is null.
    fn reverse_index(node: &Option<Box<ListNode>>) -> i32 {
        match node {
            None => 0,
            Some(n) => -1 + Self::reverse_index(&n.next),
        }
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}
