pub struct Solution {}

// Definition for singly-linked list.
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

impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }

        let mut cur = head.unwrap();
        let mut next = cur.next.take();

        loop {
            if next.is_none() {
                break;
            }

            let mut node = next.take().unwrap();
            next = node.next.take();
            let prev = std::mem::replace(&mut cur, node);
            cur.next = Some(prev);
        }

        Some(cur)
    }
}
