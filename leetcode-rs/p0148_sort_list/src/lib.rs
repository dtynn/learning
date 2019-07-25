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
    pub fn sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        Solution::sort_recursive(head)
    }

    fn sort_recursive(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }

        let mut lt_sentinel = ListNode::new(0);
        let mut lt_tail = &mut lt_sentinel;
        let mut gt_sentinel = ListNode::new(0);
        let mut gt_tail = &mut gt_sentinel;

        let mut pivot = head.unwrap();
        let mut cur_node_opt = pivot.next.take();

        loop {
            let mut cur_node = match cur_node_opt.take() {
                Some(n) => n,
                None => break,
            };

            cur_node_opt = cur_node.next.take();
            if cur_node.val < pivot.val {
                lt_tail.next.replace(cur_node);
                lt_tail = lt_tail.next.as_mut().unwrap();
            } else {
                gt_tail.next.replace(cur_node);
                gt_tail = gt_tail.next.as_mut().unwrap();
            }
        }

        let mut lt = Solution::sort_recursive(lt_sentinel.next.take());
        let gt = Solution::sort_recursive(gt_sentinel.next.take());

        pivot.next = gt;
        if lt.is_none() {
            return Some(pivot);
        }

        let mut lt_last = lt.as_mut().unwrap();
        loop {
            if lt_last.next.is_none() {
                lt_last.next.replace(pivot);
                break;
            }

            lt_last = lt_last.next.as_mut().unwrap();
        }

        lt
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn sort_list_test() {
        let mut sentinel = super::ListNode::new(0);
        let mut tail = &mut sentinel;

        let nums = vec![-1, 5, 3, 4, 0];
        for i in 0..nums.len() {
            tail.next = Some(Box::new(super::ListNode::new(nums[i])));
            tail = tail.next.as_mut().unwrap();
        }

        let sorted = super::Solution::sort_list(sentinel.next.take());
        let mut node_ref = sorted.as_ref();
        while let Some(node) = node_ref {
            println!("{}", node.val);
            node_ref = node.next.as_ref();
        }
    }
}
