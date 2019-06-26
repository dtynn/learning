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

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn build_tree(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        Solution::build_node(&preorder[..], &inorder[..])
    }

    fn build_node(preorder: &[i32], inorder: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
        assert!(preorder.len() == inorder.len());
        if preorder.len() == 0 {
            return None;
        }

        let node_val = preorder[0];
        let mut node = TreeNode::new(node_val);

        let mut node_in_inorder = 0;
        while node_in_inorder < inorder.len() {
            if inorder[node_in_inorder] == node_val {
                break;
            }

            node_in_inorder += 1;
        }

        let (inorder_left, inorder_right) = (
            &inorder[0..node_in_inorder],
            &inorder[node_in_inorder + 1..],
        );

        let left_size = inorder_left.len();
        let (preorder_left, preorder_right) =
            (&preorder[1..1 + left_size], &preorder[1 + left_size..]);

        node.left = Solution::build_node(preorder_left, inorder_left);
        node.right = Solution::build_node(preorder_right, inorder_right);

        Some(Rc::new(RefCell::new(node)))
    }
}
