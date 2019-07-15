pub struct Solution {}

// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut res = Vec::new();

        Solution::inorder_traversal_recursive(root, &mut res);

        res
    }

    fn inorder_traversal_recursive(node: Option<Rc<RefCell<TreeNode>>>, res: &mut Vec<i32>) {
        let n = match node {
            Some(n) => n,
            None => {
                return;
            }
        };

        Solution::inorder_traversal_recursive(n.as_ref().borrow().left.clone(), res);
        res.push(n.borrow().val);
        Solution::inorder_traversal_recursive(n.as_ref().borrow().right.clone(), res);
    }
}
