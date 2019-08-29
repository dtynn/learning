pub struct Solution {}

impl Solution {
    pub fn trailing_zeroes(n: i32) -> i32 {
        let mut count_of_five = 0;

        let mut divisor = 5;
        let mut divided = n / divisor;
        while divided > 0 {
            count_of_five += divided;
            divisor *= 5;
            divided = n / divisor;
        }

        count_of_five
    }
}
