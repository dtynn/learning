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
    pub fn odd_even_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }

        let mut odd_sentinel = ListNode::new(0);
        let mut odd_tail = &mut odd_sentinel;
        let mut event_sentinel = ListNode::new(0);
        let mut event_tail = &mut event_sentinel;

        let mut first = head.unwrap();
        let mut cur = first.next.take();

        odd_tail.next = Some(first);
        odd_tail = odd_tail.next.as_mut().unwrap();

        let mut node_num = 2;
        loop {
            if cur.is_none() {
                break;
            }

            let mut node = cur.take().unwrap();
            cur = node.next.take();

            if node_num % 2 == 0 {
                event_tail.next = Some(node);
                event_tail = event_tail.next.as_mut().unwrap();
            } else {
                odd_tail.next = Some(node);
                odd_tail = odd_tail.next.as_mut().unwrap();
            }

            node_num += 1;
        }

        odd_tail.next = event_sentinel.next.take();

        odd_sentinel.next.take()
    }
}
