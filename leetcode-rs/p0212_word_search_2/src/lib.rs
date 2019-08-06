//! 在回溯算法中, 通常要考虑根据目标结果集(在本题中, 即可能的words) 来进行一定的优化, 减少重复计算
pub struct Solution {}

use std::collections::{HashMap, HashSet};

struct TrieNode {
    chars: HashMap<char, TrieNode>,
    end_of_word: bool,
}

impl TrieNode {
    fn new() -> TrieNode {
        TrieNode {
            chars: HashMap::new(),
            end_of_word: false,
        }
    }
}

impl Solution {
    pub fn find_words(board: Vec<Vec<char>>, words: Vec<String>) -> Vec<String> {
        let row_size = board.len();
        if row_size == 0 {
            return vec![];
        }

        let col_size = board[0].len();

        let word_chars: Vec<Vec<char>> = words.into_iter().map(|s| s.chars().collect()).collect();

        let mut root = TrieNode::new();
        for chars in word_chars.iter() {
            Solution::add_word(&mut root, chars);
        }

        let mut res = vec![];
        let mut visited = HashSet::new();
        let mut chars = vec![];

        for row in 0..row_size {
            for col in 0..col_size {
                if root.chars.contains_key(&board[row][col]) {
                    chars.clear();
                    visited.clear();

                    Solution::find(
                        &board,
                        (row_size, col_size),
                        &mut root,
                        (row, col),
                        &mut chars,
                        &mut visited,
                        &mut res,
                    );
                }
            }
        }

        res
    }

    fn add_word(mut node_ptr: &mut TrieNode, chars: &[char]) {
        for c in chars {
            if !node_ptr.chars.contains_key(c) {
                node_ptr.chars.insert(*c, TrieNode::new());
            }

            node_ptr = node_ptr.chars.get_mut(c).unwrap();
        }

        node_ptr.end_of_word = true;
    }

    fn find(
        board: &[Vec<char>],
        (row_size, col_size): (usize, usize),
        node: &mut TrieNode,
        pos: (usize, usize),
        chars: &mut Vec<char>,
        visited: &mut HashSet<(usize, usize)>,
        res: &mut Vec<String>,
    ) {
        if visited.contains(&pos) {
            return;
        }

        let cur_char = board[pos.0][pos.1];

        chars.push(cur_char);
        visited.insert(pos);

        if let Some(mut n) = node.chars.get_mut(&cur_char) {
            if n.end_of_word {
                n.end_of_word = false;
                res.push(chars.iter().collect());
            }

            let mut next_pos_list = vec![];
            if pos.0 > 0 {
                next_pos_list.push((pos.0 - 1, pos.1));
            }

            if pos.0 < row_size - 1 {
                next_pos_list.push((pos.0 + 1, pos.1));
            }

            if pos.1 > 0 {
                next_pos_list.push((pos.0, pos.1 - 1));
            }

            if pos.1 < col_size - 1 {
                next_pos_list.push((pos.0, pos.1 + 1));
            }

            for next_pos in next_pos_list {
                Solution::find(
                    board,
                    (row_size, col_size),
                    n,
                    next_pos,
                    chars,
                    visited,
                    res,
                );
            }
        }

        chars.pop();
        visited.remove(&pos);
    }
}
