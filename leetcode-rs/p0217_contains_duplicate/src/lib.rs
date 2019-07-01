struct Solution {}

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut distincts = std::collections::HashMap::new();

        for num in nums {
            if distincts.contains_key(&num) {
                return true;
            }

            distincts.insert(num, ());
        }

        false
    }
}
