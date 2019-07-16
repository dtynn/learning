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
    pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        let mut lists = lists;
        let mut btree_m = std::collections::BTreeMap::new();
        let mut sentinel = ListNode::new(0);
        let mut cur = &mut sentinel;

        let list_size = lists.len();
        for i in 0..list_size {
            if let Some(node) = lists[i].as_ref() {
                btree_m.insert((node.val, i), ());
            }
        }

        while !btree_m.is_empty() {
            let (min_val, min_idx) = btree_m.iter().next().map(|(k, _)| (k.0, k.1)).unwrap();
            btree_m.remove(&(min_val, min_idx));

            let min_next = lists[min_idx]
                .as_mut()
                .map(|inner| inner.next.take())
                .unwrap();

            if let Some(node) = min_next.as_ref() {
                btree_m.insert((node.val, min_idx), ());
            }

            let min_node = std::mem::replace(&mut lists[min_idx], min_next);
            cur.next = min_node;
            cur = cur.next.as_mut().map(|inner| inner.as_mut()).unwrap();
        }

        sentinel.next
    }
}
