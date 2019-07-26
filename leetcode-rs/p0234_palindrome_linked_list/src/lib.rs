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
    pub fn is_palindrome(head: Option<Box<ListNode>>) -> bool {
        // 先确定 head 长度, 时间 O(n) 空间 O(1)
        // 截取尾部 n/2长度的链表并 reverse, 时间 O(n/2), 空间 O(1)
        // 对比 head 前 n/2 和 尾部链表, 得出结果 时间 O(n/2)
        unimplemented!();
    }
}
