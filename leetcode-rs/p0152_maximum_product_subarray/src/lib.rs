struct Solution {}

impl Solution {
    pub fn max_product(nums: Vec<i32>) -> i32 {
        let size = nums.len();
        if size == 0 {
            return 0;
        }

        let mut products = vec![vec![]; size];

        for i in 0..size {
            products[i].push(nums[i]);
            for j in 0..i {
                let new_product = products[j].last().unwrap() * nums[i];
                products[j].push(new_product);
            }
        }

        let mut max = std::i32::MIN;
        for row in products {
            for num in row {
                max = max.max(num);
            }
        }

        max
    }
}
