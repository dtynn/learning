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
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut sentinel_n = ListNode::new(0);
        sentinel_n.next = head;

        let mut sentinel = sentinel_n;

        let mut fast_p = Box::new(sentinel.clone());
        let mut nth_p = &mut sentinel;

        for _ in 0..n {
            fast_p = fast_p.next.as_ref().map(|n| n.clone()).unwrap();
        }

        while fast_p.next.is_some() {
            fast_p = fast_p.next.as_ref().map(|n| n.clone()).unwrap();
            nth_p = nth_p.next.as_mut().map(|n| n.as_mut()).unwrap();
        }

        let next = nth_p.next.take().and_then(|inner| inner.next);
        nth_p.next = next;

        sentinel.next
    }
}
