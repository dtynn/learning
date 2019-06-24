struct Solution {}

impl Solution {
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut result = vec![vec![]];

        let prefix = vec![];
        Solution::traverse(&mut result, &prefix, &nums[..]);

        result
    }

    fn traverse(result: &mut Vec<Vec<i32>>, prefix: &[i32], remains: &[i32]) {
        let remains_size = remains.len();
        if remains_size == 0 {
            return;
        }

        let mut i = 0;
        while i < remains_size {
            let mut new_prefix = prefix.to_owned();
            new_prefix.push(remains[i]);
            result.push(new_prefix.clone());

            let new_remains = &remains[i + 1..];
            Solution::traverse(result, &new_prefix[..], new_remains);

            i += 1;
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn p0078_test() {
        let nums = vec![1, 2, 3, 4, 5, 6];

        let res = super::Solution::subsets(nums);

        println!("{:?}", res);
    }
}
