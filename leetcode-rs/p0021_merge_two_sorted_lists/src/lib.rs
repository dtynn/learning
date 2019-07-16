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
    pub fn merge_two_lists(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut sentinel = Box::new(ListNode::new(0));
        let mut cur = &mut sentinel;
        let mut l1node = l1;
        let mut l2node = l2;

        loop {
            if l1node.is_none() {
                cur.next = l2node;
                break;
            }

            if l2node.is_none() {
                cur.next = l1node;
                break;
            }

            let l1val = l1node.as_ref().map(|inner| inner.val).unwrap();
            let l2val = l2node.as_ref().map(|inner| inner.val).unwrap();

            let next_node = if l1val < l2val {
                let mut next = l1node;
                l1node = next.as_mut().map(|inner| inner.next.take()).unwrap();
                next
            } else {
                let mut next = l2node;
                l2node = next.as_mut().map(|inner| inner.next.take()).unwrap();
                next
            };

            cur.next = next_node;
            cur = cur.next.as_mut().unwrap();
        }

        sentinel.next
    }
}
