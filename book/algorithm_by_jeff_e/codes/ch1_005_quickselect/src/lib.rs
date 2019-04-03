/// k < nums.len
pub fn quick_select(nums: &mut [i32], k: usize) -> i32 {
    let size = nums.len();
    assert!(k < size);

    if size == 1 {
        return nums[0];
    }

    let pivot = nums[0];
    let mut lt_pos = 1;
    let mut gte_pos = size;
    while lt_pos < gte_pos {
        if nums[lt_pos] < pivot {
            lt_pos += 1;
        } else {
            gte_pos -= 1;
            nums.swap(lt_pos, gte_pos);
        }
    }

    let lt_cnt = nums[1..gte_pos].len();

    if lt_cnt == k {
        return nums[0];
    }

    if lt_cnt > k {
        return quick_select(&mut nums[1..gte_pos], k);
    }

    quick_select(&mut nums[gte_pos..], k - lt_cnt - 1)
}

#[cfg(test)]
mod test {
    use super::quick_select;

    #[test]
    fn quick_select_test() {
        let cases = vec![
            (vec![1, 2, 3, 4, 5], 0, 1),
            (vec![1, 2, 3, 4, 5], 1, 2),
            (vec![1, 2, 3, 4, 5], 2, 3),
            (vec![1, 2, 3, 4, 5], 3, 4),
            (vec![1, 2, 3, 4, 5], 4, 5),
            (vec![5, 4, 3, 2, 1], 0, 1),
            (vec![5, 4, 3, 2, 1], 1, 2),
            (vec![5, 4, 3, 2, 1], 2, 3),
            (vec![5, 4, 3, 2, 1], 3, 4),
            (vec![5, 4, 3, 2, 1], 4, 5),
        ];

        for (i, c) in cases.into_iter().enumerate() {
            let mut nums = c.0.clone();
            assert_eq!(quick_select(&mut nums[..], c.1), c.2, "#{} case", i);
        }
    }
}
