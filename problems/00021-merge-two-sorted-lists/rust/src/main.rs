#![allow(dead_code)]
struct Solution {}

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
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (list1, list2) {
            (Some(l1), Some(l2)) => {
                let (cur_val, list1, list2) = if l1.val < l2.val {
                    (l1.val, l1.next, Some(l2))
                } else {
                    (l2.val, Some(l1), l2.next)
                };
                let mut node = Box::new(ListNode::new(cur_val));
                node.next = Self::merge_two_lists(list1, list2);
                Some(node)
            }
            (l1, None) if l1.is_some() => l1,
            (None, l2) if l2.is_some() => l2,
            _ => None,
        }
    }
}

fn main() {
    println!("This program doesn't have testing since there's no suitable way of complex testing");
}
